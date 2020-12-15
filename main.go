package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
	fmt.Fprint(w, "路径:"+r.URL.Path)
}
func main() {
	fmt.Println("http://localhost:3000")
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)

}
