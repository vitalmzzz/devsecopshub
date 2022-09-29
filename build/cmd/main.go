package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/vitalmzzz/devsecopshub/internal/service"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("Error initialization config %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()

	r.HandleFunc("/GetInfo", service.GetInfo).Methods("GET")
	r.HandleFunc("/GetProject", service.GetProject).Methods("GET")
	r.HandleFunc("/GetAppSecHub", service.GetAppSecHub).Methods("GET")
	r.HandleFunc("/GetJira", service.GetJira).Methods("GET")

	err := http.ListenAndServe(viper.GetString("host_string"), r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
