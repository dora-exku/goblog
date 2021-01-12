package controllers

import (
	"fmt"
	"goblog/app/models/user"
	"goblog/pkg/view"
	"net/http"
)

type AuthController struct {
}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {

	view.RenderSimple(w, view.D{}, "auth.register")
}

func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	_user := user.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	_user.Create()
	if _user.ID > 0 {
		fmt.Fprint(w, "用户创建成功ID："+_user.GetStringId())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "用户创建失败")
	}

}
