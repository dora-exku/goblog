package session

import (
	"goblog/pkg/config"
	"goblog/pkg/logger"
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte(config.GetString("key")))

var Session *sessions.Session

var Request *http.Request

var Response http.ResponseWriter

func SessionStart(w http.ResponseWriter, r *http.Request) {
	var err error
	Session, err = Store.Get(r, config.GetString("session_name"))
	logger.LogError(err)
	Request = r
	Response = w
}

func Put(name string, value interface{}) {
	Session.Values[name] = value
	Save()
}

func Get(name string) interface{} {
	return Session.Values[name]
}

func Forget(name string) {
	delete(Session.Values, name)
	Save()
}

func Flush() {
	Session.Options.MaxAge = -1
	Save()
}

func Save() {
	err := Session.Save(Request, Response)
	logger.LogError(err)
}
