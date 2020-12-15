package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		fmt.Fprint(w, "这里是首页")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "这里是关于页面")
	} else {
		fmt.Fprint(w, "页面未找到")
	}
}

func main() {
	fmt.Println("http://localhost:3000")
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)

}
