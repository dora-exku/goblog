package middlewares

import (
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"net/http"
)

func Guest(next HttpHandlerFunc) HttpHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if auth.Check() {

			flash.Warning("登录用户不允许访问当前页面")
			http.Redirect(w, r, "/", http.StatusFound)
			return

		}
		next(w, r)
	}
}
