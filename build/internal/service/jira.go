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

//Получаем обновления по задачам Jira для проекта SIBDSO
func GetJira(w http.ResponseWriter, r *http.Request) {

	jira_token := os.Getenv("JIRA_TOKEN")

	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006 15:04:05"), "Jira: Receive project updates SIBDSO")

	client := &http.Client{}

	// var startAt = [6]int {}
	// startAt[0] = 0
	// startAt[1] = 3000
	// startAt[2] = 6000
	// startAt[3] = 9000
	// startAt[4] = 12000
	// startAt[5] = 15000

	req, err := http.NewRequest("GET", "https://server/jira/rest/api/2/search?jql=project=SIBDSO+order+by+updated&fields=id,summary,key,labels,components,assignee,status,created,resolutiondate,updated,aggregateprogress&maxResults=3000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// req, err := http.NewRequest("GET", "https://server/jira/rest/api/2/search?jql=project=SIBDSO+order+by+created&fields=id,summary,key,labels,assignee,status,created,updated,aggregateprogress&startAt=9000&maxResults=3000", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+string(jira_token))

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
	fmt.Println("Parsing response body")
	fmt.Println(string(body))

	var data repository.Jira

	json.Unmarshal(body, &data)
	AddJira(data)
	GetWorkTime(data)

	// fmt.Println("ТЕСТ")
	// for k := range data.Issues {
	// 	fmt.Println(data.Issues[k].Key)
	// }

	// for k := range data.Issues {
	// 	for v := range data.Issues[k].Fields.Components {
	// 		fmt.Println(data.Issues[k].Fields.Components[v].Name)
	// 	}
	// }

	w.Write([]byte("Completed"))
}

//Получаем информацию по учету рабочего времени сотрудников
func GetWorkTime(data repository.Jira) {

	jira_token := os.Getenv("JIRA_TOKEN")

	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006 15:04:05"), "Jira: Get Time Tracker")

	client_2 := &http.Client{}

	for i, _ := range data.Issues {
		req_2, err := http.NewRequest("GET", fmt.Sprintf("https://server/jira/rest/api/2/issue/%v/worklog", data.Issues[i].Key), nil)

		req_2.Header.Add("Content-Type", "application/json")
		req_2.Header.Add("Authorization", "Basic "+string(jira_token))
		if err != nil {
			fmt.Println(err)
		}

		res, err := client_2.Do(req_2)
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

		var workTime repository.WorkTime

		json.Unmarshal(body, &workTime)
		AddWorkTime(workTime)

		// fmt.Println("ТЕСТ")
		// for k := range workTime.Worklogs {
		// 	fmt.Println(workTime.Worklogs[k].Comment)
		// }

	}
}

//Пишем в базу данных информацию по задачам
func AddJira(data repository.Jira) {

	connStr := os.Getenv("CONNSTR")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// for i, _ := range data.Issues {
	// 	defer db.Close()
	// 	_, err = db.Exec("insert into jira_report (id, key, summary, name, status, labels, created, updated, time_taken) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
	// 		data.Issues[i].ID, data.Issues[i].Key, data.Issues[i].Fields.Summary, data.Issues[i].Fields.Assignee.DisplayName, data.Issues[i].Fields.Status.Name, data.Issues[i].Fields.Labels, data.Issues[i].Fields.Created, data.Issues[i].Fields.Updated, data.Issues[i].Fields.Aggregateprogress.Total)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//  }

	for i, _ := range data.Issues {
		defer db.Close()
		_, err = db.Exec(`insert into jira_report_v2 (id, key, summary, name, status, labels, created, resolutiondate, updated, time_taken, url)
	 	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) on conflict (id) do update set
	 	key=$2, summary=$3, name=$4, status=$5, labels=$6, created=$7, resolutiondate=$8, updated=$9, time_taken=$10, url=$11`,
			data.Issues[i].ID, data.Issues[i].Key, data.Issues[i].Fields.Summary, data.Issues[i].Fields.Assignee.DisplayName, data.Issues[i].Fields.Status.Name, data.Issues[i].Fields.Labels, data.Issues[i].Fields.Created, data.Issues[i].Fields.Resolutiondate, data.Issues[i].Fields.Updated, data.Issues[i].Fields.Aggregateprogress.Total, fmt.Sprintf("https://server/jira/browse/%v", data.Issues[i].Key))
		if err != nil {
			panic(err)
		}
	}

	for i, _ := range data.Issues {
		for v, _ := range data.Issues[i].Fields.Components {
			defer db.Close()
			_, err = db.Exec(`insert into jira_report_v3 (id, key, summary, name, status, labels, components, created, resolutiondate, updated, time_taken, url)
 		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) on conflict (id) do update set
 		key=$2, summary=$3, name=$4, status=$5, labels=$6, components=$7, created=$8, resolutiondate=$9, updated=$10, time_taken=$11, url=$12`,
				data.Issues[i].ID, data.Issues[i].Key, data.Issues[i].Fields.Summary, data.Issues[i].Fields.Assignee.DisplayName, data.Issues[i].Fields.Status.Name, data.Issues[i].Fields.Labels, data.Issues[i].Fields.Components[v].Name, data.Issues[i].Fields.Created, data.Issues[i].Fields.Resolutiondate, data.Issues[i].Fields.Updated, data.Issues[i].Fields.Aggregateprogress.Total, fmt.Sprintf("https://server/jira/browse/%v", data.Issues[i].Key))
			if err != nil {
				panic(err)
			}
		}
	}

}

//Пишем в базу данных информацию списанию рабочих часов
func AddWorkTime(workTime repository.WorkTime) {

	connStr := os.Getenv("CONNSTR")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	for i, _ := range workTime.Worklogs {
		defer db.Close()
		_, err = db.Exec(`insert into jira_work_time (id, issue_id, name, comment, created, updated, started, time_spent, time_spent_seconds) 
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9) on conflict (id) do update set
			issue_id=$2, name=$3, comment=$4, created=$5, updated=$6, started=$7, time_spent=$8, time_spent_seconds=$9`,
			workTime.Worklogs[i].ID, workTime.Worklogs[i].IssueID, workTime.Worklogs[i].Author.DisplayName, workTime.Worklogs[i].Comment, workTime.Worklogs[i].Created, workTime.Worklogs[i].Updated, workTime.Worklogs[i].Started, workTime.Worklogs[i].TimeSpent, workTime.Worklogs[i].TimeSpentSeconds)
		if err != nil {
			panic(err)
		}
	}

}
