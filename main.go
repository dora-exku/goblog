package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "这里是首页 <a href=\"/about\">关于</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 NOT FOUND")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "这里是关于页面 <a href=\"/\">首页</a>")
}

func main() {
	fmt.Println("http://localhost:3000")
	router := http.NewServeMux()

	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":3000", router)

}
