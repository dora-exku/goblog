package controllers

import (
	"database/sql"
	"fmt"
	"goblog/app/models/article"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"html/template"
	"net/http"
	"strconv"
	"unicode/utf8"

	"gorm.io/gorm"
)

type ArticlesController struct{}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	article, err := article.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "文章不存在")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误")
		}
	} else {
		tmpl, err := template.New("show.gohtml").
			Funcs(template.FuncMap{
				"RouteName2Url": route.Name2URL,
				"Int64ToString": types.Int64ToString,
			}).
			ParseFiles("resources/views/articles/show.gohtml")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, article)
	}
}

func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "文章列表")
	articles, err := article.GetAll()

	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 server error")
	} else {
		tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
		logger.LogError(err)
		tmpl.Execute(w, articles)
	}
}

func validateArticleFormData(title string, content string) map[string]string {
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
	return errors
}

func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")

	errors := validateArticleFormData(title, content)

	if len(errors) == 0 {
		_article := article.Article{
			Title:   title,
			Content: content,
		}
		_article.Create()
		if _article.ID > 0 {
			fmt.Fprint(w, "信息加入成功ID:"+strconv.FormatInt(_article.ID, 10))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误")
		}
	} else {

		storeUrl := route.Name2URL("articles.store")

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

type ArticlesFormData struct {
	Title, Content string
	URL            string
	Errors         map[string]string
}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {

	storeUrl := route.Name2URL("articles.store")

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

func (*ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	// 查询一条数据
	article, err := article.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "文章不存在")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误")
		}
	} else {

		updateUrl := route.Name2URL("articles.show", "id", id)
		data := ArticlesFormData{
			Title:   article.Title,
			Content: article.Content,
			URL:     updateUrl,
			Errors:  nil,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
		logger.LogError(err)
		tmpl.Execute(w, data)
	}
}

func (*ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	// 查询一条数据
	_article, err := article.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "文章不存在")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误")
		}
	} else {
		title := r.PostFormValue("title")
		content := r.PostFormValue("content")
		errors := validateArticleFormData(title, content)

		if len(errors) == 0 {
			_article.Title = title
			_article.Content = content
			rowsAffected, err := _article.Update()
			if err != nil {
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "服务器内部错误")
			}
			if rowsAffected > 0 {
				fmt.Fprint(w, "信息修改成功ID:"+strconv.FormatInt(_article.ID, 10))
			} else {
				fmt.Fprint(w, "您没有做出任何更改")
			}
		} else {

			updateUrl := route.Name2URL("articles.update", "id", id)
			data := ArticlesFormData{
				Title:   title,
				Content: content,
				URL:     updateUrl,
				Errors:  errors,
			}
			tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
			logger.LogError(err)
			tmpl.Execute(w, data)
		}
	}
}
