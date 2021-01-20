package controllers

import (
	"database/sql"
	"fmt"
	"goblog/app/models/article"
	"goblog/app/requests"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"

	"gorm.io/gorm"
)

type ArticlesController struct{}

type ArticlesFormData struct {
	Title, Content string
	Article        article.Article
	URL            string
	Errors         map[string]string
}

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

		view.Render(w, view.D{
			"Article": article,
		}, "articles.show", "articles._article_meta")
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
		view.Render(w, view.D{
			"Articles": articles,
		}, "articles.index", "articles._article_meta")
	}
}

func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")

	_article := article.Article{
		Title:   title,
		Content: content,
	}

	errors := requests.ValidateArticleForm(_article)

	if len(errors) == 0 {

		_article.Create()
		if _article.ID > 0 {
			flash.Success("文章添加成功")
			http.Redirect(w, r, route.Name2URL("articles.show", "id", _article.GetStringId()), http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误")
		}
	} else {
		view.Render(w, view.D{
			"Article": _article,
			"Errors":  errors,
		}, "articles.create", "articles._form_field")
	}
}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
}

func (*ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
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
		view.Render(w, view.D{
			"Title":   _article.Title,
			"Content": _article.Content,
			"Article": _article,
			"Errors":  nil,
		}, "articles.edit", "articles._form_field")
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
		_article.Title = title
		_article.Content = content
		errors := requests.ValidateArticleForm(_article)

		if len(errors) == 0 {

			rowsAffected, err := _article.Update()
			if err != nil {
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "服务器内部错误")
			}
			if rowsAffected > 0 {
				fmt.Fprint(w, "信息修改成功ID:"+_article.GetStringId())
			} else {
				fmt.Fprint(w, "您没有做出任何更改")
			}
		} else {

			view.Render(w, view.D{
				"Title":   title,
				"Content": content,
				"Article": _article,
				"Errors":  errors,
			}, "articles.edit", "articles._form_field")
		}
	}
}

func (*ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	article, err := article.Get(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "数据不存在")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误")
		}
	} else {
		rowsAffected, err := article.Delete()

		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误")
		} else {
			if rowsAffected > 0 {
				indexUrl := route.Name2URL("articles.index")
				http.Redirect(w, r, indexUrl, http.StatusFound)
			} else {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "文章未找到")
			}
		}
	}
}
