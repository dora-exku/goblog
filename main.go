package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "这里是首页 <a href=\"/about\">关于</a>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "这里是关于页面 <a href=\"/\">首页</a>")
	} else {
		fmt.Fprint(w, "页面未找到")
	}
}

func main() {
	fmt.Println("http://localhost:3000")
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)

}
