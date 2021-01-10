package controllers

import (
	"database/sql"
	"fmt"
	"goblog/app/models/article"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
	"unicode/utf8"

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

		view.Render(w, article, "articles.show")
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
		view.Render(w, articles, "articles.index")
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
			fmt.Fprint(w, "信息加入成功ID:"+_article.GetStringId())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器内部错误")
		}
	} else {
		view.Render(w, ArticlesFormData{
			Title:   title,
			Content: content,
			Errors:  errors,
		}, "articles.create", "articles._form_field")
	}
}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, ArticlesFormData{}, "articles.create", "articles._form_field")
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
		view.Render(w, ArticlesFormData{
			Title:   _article.Title,
			Content: _article.Content,
			Article: _article,
			Errors:  nil,
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
				fmt.Fprint(w, "信息修改成功ID:"+_article.GetStringId())
			} else {
				fmt.Fprint(w, "您没有做出任何更改")
			}
		} else {

			view.Render(w, ArticlesFormData{
				Title:   title,
				Content: content,
				Article: _article,
				Errors:  errors,
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
