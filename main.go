package main

import (
	"log"
	"net/http"
	"os"

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

	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggingMiddleware)
	router.Use(jsonMiddleware)

	router.HandleFunc("/api/posts", AllPostsEndpoint).Methods("GET")
	router.HandleFunc("/api/posts", CreatePostEndpoint).Methods("POST")

	log.Fatal(http.ListenAndServe(config.ServerPort, router))
}
