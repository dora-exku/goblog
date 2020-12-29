package main

import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"net/http"
)

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetopRoute()
	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
