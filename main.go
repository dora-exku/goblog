package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"time"
	"unicode/utf8"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()
var db *sql.DB

type ArticlesFormData struct {
	Title, Content string
	URL            *url.URL
	Errors         map[string]string
}

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
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")

	errors := make(map[string]string)

	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) >= 40 {
		errors["title"] = "标题长度限制为 3-40"
	}

	if content == "" {
		errors["content"] = "内容不能为空"
	} else if utf8.RuneCountInString(content) < 10 {
		errors["content"] = "内容长度需要大于10"
	}

	if len(errors) == 0 {
		fmt.Fprint(w, "验证通过<br>")
		fmt.Fprintf(w, "标题内容 %s<br>", title)
		fmt.Fprintf(w, "提交内容 %s<br>", content)
	} else {

		storeUrl, _ := router.Get("articles.store").URL()

		data := ArticlesFormData{
			Title:   title,
			Content: content,
			URL:     storeUrl,
			Errors:  errors,
		}

		tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, data)
	}
}

func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {

	storeUrl, _ := router.Get("articles.store").URL()

	data := ArticlesFormData{
		Title:   "",
		Content: "",
		URL:     storeUrl,
		Errors:  nil,
	}
	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")

	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, data)
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

func initDB() {
	var err error

	config := mysql.Config{
		User:                 "homestead",
		Passwd:               "secret",
		Addr:                 "127.0.0.1:33060",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	checkErr(err)

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)

	err = db.Ping()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB()
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
