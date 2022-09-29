package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/vitalmzzz/devsecopshub/internal/repository"
)

//Получаем информацию по проектам
func GetAppSecHub(w http.ResponseWriter, r *http.Request) {

	appsechub_token := os.Getenv("APPSECHUB_TOKEN")

	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006 15:04:05"), "AppSecHub: Get project ID")

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://server/hub/rest/application/v2?pageIndex=0&pageSize=100", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Authorization", "Token "+string(appsechub_token))

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var appInfo repository.RestApplication

	json.Unmarshal(body, &appInfo)
	DefectsInfo(appInfo)
	CodebaseInfo(appInfo)
	AddProjectInfo(appInfo)
	ArtifactInfo(appInfo)

	w.Write([]byte("Completed"))

}

//Получаем информацию по артифактам
func ArtifactInfo(appInfo repository.RestApplication) {

	appsechub_token := os.Getenv("APPSECHUB_TOKEN")

	dt := time.Now()

	fmt.Println(dt.Format("01-02-2006 15:04:05"), "AppSecHub: Get information on artifacts")
	client_1 := &http.Client{}

	for i, _ := range appInfo.FilteredEntities {
		req_1, err := http.NewRequest("GET", fmt.Sprintf("https://server/hub/rest/application/%v/artifact", appInfo.FilteredEntities[i].ID), nil)

		req_1.Header.Add("Content-Type", "application/json")
		req_1.Header.Add("X-Authorization", "Token "+string(appsechub_token))
		if err != nil {
			fmt.Println(err)
		}

		resp, err := client_1.Do(req_1)
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		var artifacts repository.Artifacts

		json.Unmarshal(bodyBytes, &artifacts)
		AddArtifactsInfo(artifacts)
	}
}

//Получем информацию по дефектам
func DefectsInfo(appInfo repository.RestApplication) {

	appsechub_token := os.Getenv("APPSECHUB_TOKEN")

	dt := time.Now()

	fmt.Println(dt.Format("01-02-2006 15:04:05"), "AppSecHub: Get information about defects")

	client_2 := &http.Client{}

	for i, _ := range appInfo.FilteredEntities {
		req_2, err := http.NewRequest("GET", fmt.Sprintf("https://server/hub/rest/defects?sort=1&appId=%v&pageIndex=0&pageSize=14", appInfo.FilteredEntities[i].ID), nil)

		req_2.Header.Add("Content-Type", "application/json")
		req_2.Header.Add("X-Authorization", "Token "+string(appsechub_token))
		if err != nil {
			fmt.Println(err)
		}

		resp, err := client_2.Do(req_2)
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		var defects repository.Defects

		json.Unmarshal(bodyBytes, &defects)
		AddDefectsInfo(defects)
	}
}

//Получаем информацию по кодовой базе
func CodebaseInfo(appInfo repository.RestApplication) {

	appsechub_token := os.Getenv("APPSECHUB_TOKEN")

	dt := time.Now()

	fmt.Println(dt.Format("01-02-2006 15:04:05"), "AppSecHub: Get information on codebase")

	client_3 := &http.Client{}

	for i, _ := range appInfo.FilteredEntities {
		req_3, err := http.NewRequest("GET", fmt.Sprintf("https://server/hub/rest/application/%v/codebase", appInfo.FilteredEntities[i].ID), nil)

		req_3.Header.Add("Content-Type", "application/json")
		req_3.Header.Add("X-Authorization", "Token "+string(appsechub_token))
		if err != nil {
			fmt.Println(err)
		}

		resp, err := client_3.Do(req_3)
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		var codebase repository.Codebase

		json.Unmarshal(bodyBytes, &codebase)
		AddCodebaseInfo(codebase)
	}
}

//Пишем в базу данных информацию по проектам
func AddProjectInfo(appInfo repository.RestApplication) {

	connStr := os.Getenv("CONNSTR")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	for i, _ := range appInfo.FilteredEntities {
		defer db.Close()
		_, err = db.Exec(`insert into appsechub_projects (id, name, code, defects_total, issues_total, ssdl_total, codebase_size, project_link) 
			values ($1, $2, $3, $4, $5, $6, $7, $8) on conflict (id) do update set
			name=$2, code=$3, defects_total=$4, issues_total=$5, ssdl_total=$6, codebase_size=$7, project_link=$8`,
			appInfo.FilteredEntities[i].ID, appInfo.FilteredEntities[i].Name, appInfo.FilteredEntities[i].Code, appInfo.FilteredEntities[i].Summary.Defects.Total, appInfo.FilteredEntities[i].Summary.Issues.Total, appInfo.FilteredEntities[i].Summary.SsdlTasksSummary.SsdlTotal, appInfo.FilteredEntities[i].Summary.CodebaseSize, fmt.Sprintf("https://server/#/appprofile/%v/info", appInfo.FilteredEntities[i].ID))
		if err != nil {
			panic(err)
		}
	}
}

//Пишем в базу данных информацию по артифактам
func AddArtifactsInfo(artifacts repository.Artifacts) {

	connStr := os.Getenv("CONNSTR")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	for i, _ := range artifacts {
		defer db.Close()
		_, err = db.Exec(`insert into appsechub_artifacts (id, repository_url, artifact ) 
			values ($1, $2, $3) on conflict (id) do update set
			repository_url=$2, artifact=$3`,
			artifacts[i].AppID, artifacts[i].ArtifactRepository.URL, artifacts[i].Properties.DockerArtifact)
		if err != nil {
			panic(err)
		}
	}
}

//Пишем в базу данных информацию по кодовой базе
func AddCodebaseInfo(codebase repository.Codebase) {

	connStr := os.Getenv("CONNSTR")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	for i, _ := range codebase {
		defer db.Close()
		_, err = db.Exec(`insert into appsechub_codebase (id, name, link, branch, active, appId) 
			values ($1, $2, $3, $4, $5, $6) on conflict (id) do update set
			name=$2, link=$3, branch=$4, active=$5, appId=$6`,
			codebase[i].ID, codebase[i].Name, codebase[i].Link, codebase[i].Branch, codebase[i].Active, codebase[i].AppID)
		if err != nil {
			panic(err)
		}
	}
}

//Пишем в базу данных информацию по дефектам
func AddDefectsInfo(defects repository.Defects) {

	connStr := os.Getenv("CONNSTR")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	for k := range defects.FilteredEntities {
		for v := range defects.FilteredEntities[k].ExternalBugs {
			defer db.Close()
			_, err = db.Exec(`insert into appsechub_defects (id, appid, summary, priority_id, jira_link, status, created) 
			values ($1, $2, $3, $4, $5, $6, $7) on conflict (id) do update set
			appid=$2, summary=$3, priority_id=$4, jira_link=$5, status=$6, created=$7`,
				defects.FilteredEntities[k].ID, defects.FilteredEntities[k].AppID, defects.FilteredEntities[k].Summary, defects.FilteredEntities[k].PriorityID, defects.FilteredEntities[k].ExternalBugs[v].Link, defects.FilteredEntities[k].Status, defects.FilteredEntities[k].CreateTs)
			if err != nil {
				panic(err)
			}
		}
	}
}
