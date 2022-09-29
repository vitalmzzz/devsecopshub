package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/vitalmzzz/devsecopshub/internal/repository"
)

var wg sync.WaitGroup

//Получаем информацию по проектам
func GetProject(w http.ResponseWriter, r *http.Request) {

	gitlab_token := os.Getenv("GITLAB_TOKEN")

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://server/api/v4/projects", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(gitlab_token))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var projects []repository.GitProjects

	json.Unmarshal(body, &projects)
	Add_GetProject(projects)

	if LastPage(res) != 0 {
		for i := 2; i <= LastPage(res); i++ {
			wg.Add(1)
			go GetNextPage(i)
		}
	}

	wg.Wait()
	w.Write([]byte("Completed"))
}

//Считаем количество страниц
func LastPage(resp *http.Response) int {
	page := 0
	if links := resp.Header.Values("link"); len(links) > 0 {
		for _, link := range strings.Split(links[0], ",") {
			segments := strings.Split(strings.TrimSpace(link), ";")
			if len(segments) < 2 {
				continue
			}
			if !strings.HasPrefix(segments[0], "<") || !strings.HasSuffix(segments[0], ">") {
				continue
			}
			urls, err := url.Parse(segments[0][1 : len(segments[0])-1])
			if err != nil {
				continue
			}
			m, _ := url.ParseQuery(urls.RawQuery)
			page, _ = strconv.Atoi(m["page"][0])
		}
	}
	return page
}

//Выполняем обработку данных постранично
func GetNextPage(page int) {

	gitlab_token := os.Getenv("GITLAB_TOKEN")

	defer wg.Done()

	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006 15:04:05"), "GitLab: Completing page processing", page)

	client := &http.Client{}

	var r []repository.GitProjects

	req, err := http.NewRequest("GET", fmt.Sprintf("https://server/api/v4/projects?&page_size=100&page=%v", page), nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(gitlab_token))
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(bodyBytes, &r)
	Add_GetProject(r)
}

//Пишем в базу данных
func Add_GetProject(projects []repository.GitProjects) {

	connStr := os.Getenv("CONNSTR")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	for i, _ := range projects {
		defer db.Close()
		_, err = db.Exec(`insert into gitlab_projects (id, name, description, path_with_namespace, web_url, archived, username, owner_name, namespace_id, namespace_name) 
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) on conflict (id) do update set
			name=$2, description=$3, path_with_namespace=$4, web_url=$5, archived=$6, username=$7, owner_name=$8, namespace_id=$9, namespace_name=$10`,
			projects[i].Id, projects[i].Name, projects[i].Description, projects[i].Path_with_namespace, projects[i].Web_url, projects[i].Archived, projects[i].Owner.Username, projects[i].Owner.Name, projects[i].Namespace.Id, projects[i].Namespace.Name)
		if err != nil {
			panic(err)
		}
	}
}
