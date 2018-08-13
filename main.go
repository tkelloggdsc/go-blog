package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var dao = DAO{}
var config = Config{}

func init() {
	environment := os.Getenv("APP_ENV")
	if environment == "" {
		log.Fatal("You must set an APP_ENV variable indicating to start the application in 'dev', 'prod', or 'test'")
	}

	config.Read(environment)

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	println("Server running on:" + config.ServerPort)

	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1/").Subrouter()
	api.HandleFunc("/posts", AllPostsEndpoint).Methods("GET")
	api.HandleFunc("/posts", CreatePostEndpoint).Methods("POST")

	router.PathPrefix("/static").Handler(http.FileServer(http.Dir(config.StaticAssetsPath)))
	router.PathPrefix("/").HandlerFunc(IndexHandler(config.StaticEntryPath))

	server := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		Addr:         "127.0.0.1" + config.ServerPort,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
