package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var dao = DAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	println("server up")

	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggingMiddleware)
	router.Use(jsonMiddleware)

	router.HandleFunc("/api/posts", PostsIndex)

	http.ListenAndServe(config.ServerPort, router)
}
