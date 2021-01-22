package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/models/user"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

type UsersController struct {
	BaseController
}

func (*UsersController) Index() {

}

func (us *UsersController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_user, err := user.Get(id)

	if err != nil {
		us.ResponseForSQLError(w, err)
	} else {
		articles, PagerData, err := article.GetByUserId(_user.GetStringId(), r, 2)
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器错误")
		} else {
			view.Render(w, view.D{
				"Articles":  articles,
				"PagerData": PagerData,
			}, "articles.index", "articles._article_meta")
		}
	}
}

func (*UsersController) Edit() {

}

func (*UsersController) Update() {

}
