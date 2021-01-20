package middlewares

import (
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"net/http"
)

func Auth(next HttpHandlerFunc) HttpHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if !auth.Check() {

			flash.Warning("您还未登录，请登录")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next(w, r)
	}
}
