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
	"github.com/vitalmzzz/devsecopshub/internal/keycloak"
	"github.com/vitalmzzz/devsecopshub/internal/repository"
)

//Получаем информацию по внутренним и внешним системам
func GetInfo(w http.ResponseWriter, r *http.Request) {

	Token := keycloak.GetToken()

	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006 15:04:05"), "App.Farm: Get a list of all internal and external systems")

	//Получаем все внутренние системы
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://server/api/v2/environments/dev/information-systems", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+string(Token.Access_token))

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

	var internal repository.InternalSystem

	json.Unmarshal(body, &internal)
	InternalInfo(internal)

	//Получаем все внешние системы
	client_1 := &http.Client{}

	req_1, err := http.NewRequest("GET", "https://server/api/v2/environments/dev/external-systems", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req_1.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req_1.Header.Add("Authorization", "Bearer "+string(Token.Access_token))

	res_1, err := client_1.Do(req_1)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res_1.Body.Close()

	body_1, err := ioutil.ReadAll(res_1.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	var external repository.ExternalSystems

	json.Unmarshal(body_1, &external)
	ExternalInfo(external)

	wg.Wait()
	w.Write([]byte("Completed"))
}

//Получаем информацию по внутренним сервисам
func InternalInfo(internal repository.InternalSystem) {

	Token := keycloak.GetToken()

	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006 15:04:05"), "App.Farm: Get information about internal services")

	client_2 := &http.Client{}

	for i, _ := range internal {
		req_2, err := http.NewRequest("GET", fmt.Sprintf("https://server/api/v2/environments/dev/information-systems/%v/services", internal[i].ID), nil)

		req_2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req_2.Header.Add("Authorization", "Bearer "+string(Token.Access_token))
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

		var intService repository.InternalServices

		json.Unmarshal(body, &intService)
		AddInternalInfo(intService)

	}
}

//Получаем информацию по внешним сервисам
func ExternalInfo(external repository.ExternalSystems) {

	Token := keycloak.GetToken()

	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006 15:04:05"), "App.Farm: Get information about external services")

	client_3 := &http.Client{}

	for i, _ := range external {
		req_3, err := http.NewRequest("GET", fmt.Sprintf("https://server/api/v2/environments/dev/external-systems/%v/services", external[i].ID), nil)

		req_3.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req_3.Header.Add("Authorization", "Bearer "+string(Token.Access_token))
		if err != nil {
			fmt.Println(err)
		}

		res, err := client_3.Do(req_3)
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

		var extService repository.ExternalServices

		json.Unmarshal(body, &extService)
		AddExternalInfo(extService)
	}
}

//Пишем в базу данных информацию по внутренним сервисам
func AddInternalInfo(intService repository.InternalServices) {

	connStr := os.Getenv("CONNSTR")
	dt := time.Now()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// for i, _ := range intService {
	// 	defer db.Close()
	// 	_, err = db.Exec("insert into appfarm_services (id, environment, information_system_id, name, description, project_path, service_type, owner, portal_link, update) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
	// 		intService[i].ID, intService[i].Environment, intService[i].InformationSystemID, intService[i].Name, intService[i].Description, intService[i].ProjectPath, intService[i].ServiceType, intService[i].System.Owner, fmt.Sprintf("https://server/environments/dev/systems/internal/%v", intService[i].InformationSystemID), fmt.Sprintf(dt.Format("01-02-2006 15:04:05")))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	for i, _ := range intService {
		defer db.Close()
		_, err = db.Exec(`insert into appfarm_services (id, environment, information_system_id, name, description, project_path, service_type, owner, portal_link, update, internal_url, public_url) 
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) on conflict (id) do update set
			environment=$2, information_system_id=$3, name=$4, description=$5, project_path=$6, service_type=$7, owner=$8, portal_link=$9, update=$10, internal_url=$11, public_url=$12`,
			intService[i].ID, intService[i].Environment, intService[i].InformationSystemID, intService[i].Name, intService[i].Description, intService[i].ProjectPath, intService[i].ServiceType, intService[i].System.Owner, fmt.Sprintf("https://server/environments/dev/systems/internal/%v", intService[i].InformationSystemID), fmt.Sprintf(dt.Format("01-02-2006 15:04:05")), intService[i].Addresses.InternalURL, intService[i].Addresses.PublicURL)
		if err != nil {
			panic(err)
		}
	}
}

//Пишем в базу данных информацию по внешним сервисам
func AddExternalInfo(extService repository.ExternalServices) {

	connStr := os.Getenv("CONNSTR")
	dt := time.Now()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// for i, _ := range extService {
	// 	defer db.Close()
	// 	_, err = db.Exec("insert into appfarm_services (id, environment, information_system_id, name, description, project_path, service_type, owner, portal_link, update) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
	// 		extService[i].ID, extService[i].Environment, extService[i].InformationSystemID, extService[i].Name, extService[i].Description, extService[i].ProjectPath, extService[i].ServiceType, extService[i].System.Owner, fmt.Sprintf("https://server
	/environments/dev/systems/external/%v", extService[i].InformationSystemID), fmt.Sprintf(dt.Format("01-02-2006 15:04:05")))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	for i, _ := range extService {
		defer db.Close()
		_, err = db.Exec(`insert into appfarm_services (id, environment, information_system_id, name, description, project_path, service_type, owner, portal_link, update, internal_url, public_url) 
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) on conflict (id) do update set
			environment=$2, information_system_id=$3, name=$4, description=$5, project_path=$6, service_type=$7, owner=$8, portal_link=$9, update=$10, internal_url=$11, public_url=$12`,
			extService[i].ID, extService[i].Environment, extService[i].InformationSystemID, extService[i].Name, extService[i].Description, extService[i].ProjectPath, extService[i].ServiceType, extService[i].System.Owner, fmt.Sprintf("https://server/environments/dev/systems/external/%v", extService[i].InformationSystemID), fmt.Sprintf(dt.Format("01-02-2006 15:04:05")), extService[i].Addresses.InternalURL, extService[i].Addresses.PublicURL)
		if err != nil {
			panic(err)
		}
	}
}
