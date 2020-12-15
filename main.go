package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>欢迎来到Blog</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "这里是关于页面 <a href=\"/\">首页</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "404 not found")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "文章ID:"+id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "文章列表")
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "添加文章")
}

func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
	<head>
	<title>文章添加</title>
	</head>
	<body>
	<form action="%s" method="POST">
	<p><input type="text" name="title"></p>
	<p><textarea name="content" row="4"></textarea></p>
	<p><button>提交</button></p>
	</form>
	</body>
	</html>
	`

	storeUrl, _ := router.Get("articles.store").URL()

	fmt.Fprintf(w, html, storeUrl)
}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("http://localhost:3000")

	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")
	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", articlesCreateHandler).Methods("GET").Name("articles.create")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	homeUrl, _ := router.Get("home").URL()
	fmt.Println("home url:", homeUrl)
	articleUrl, _ := router.Get("articles.show").URL("id", "23")
	fmt.Println("articles url:", articleUrl)

	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3000", removeTrailingSlash(router))

}
