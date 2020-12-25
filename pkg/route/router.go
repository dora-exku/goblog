package route

import (
	"github.com/gorilla/mux"
)

var Router *mux.Router

func Initialize() {
	Router = mux.NewRouter()
}

// 通过路由名称获取url
func Name2URL(name string, pairs ...string) string {
	url, err := Router.Get(name).URL(pairs...)
	if err != nil {
		//checkError(err)
		return ""
	}
	return url.String()
}
