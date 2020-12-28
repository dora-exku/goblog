package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

var route *mux.Router

func SetRoute(r *mux.Router) {
	route = r
}

// 通过路由名称获取url
func Name2URL(name string, pairs ...string) string {
	url, err := route.Get(name).URL(pairs...)
	if err != nil {
		//checkError(err)
		return ""
	}
	return url.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
