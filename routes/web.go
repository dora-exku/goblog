package routes

import (
	"goblog/app/http/controllers"
	"goblog/app/http/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterWebRoutes(r *mux.Router) {

	pc := new(controllers.PagesController)

	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	// r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	ac := new(controllers.ArticlesController)
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/", ac.Index).Methods("GET").Name("articles.index")

	r.HandleFunc("/articles", middlewares.Auth(ac.Store)).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/create", middlewares.Auth(ac.Create)).Methods("GET").Name("articles.create")

	r.HandleFunc("/articles/{id:[0-9]+}/edit", middlewares.Auth(ac.Edit)).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}", middlewares.Auth(ac.Update)).Methods("POST").Name("articles.update")
	r.HandleFunc("/articles/{id:[0-9]+}/delete", middlewares.Auth(ac.Delete)).Methods("POST").Name("articles.delete")

	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// r.Use(middlewares.ForceHTML)

	// 登陆注册
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", middlewares.Guest(auc.Register)).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", middlewares.Guest(auc.DoRegister)).Methods("POST").Name("auth.doregister")

	r.HandleFunc("/auth/login", middlewares.Guest(auc.Login)).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin", middlewares.Guest(auc.DoLogin)).Methods("POST").Name("auth.dologin")
	r.HandleFunc("/auth/logout", middlewares.Auth(auc.Logout)).Methods("POST").Name("auth.logout")

	// 用户
	user := new(controllers.UsersController)
	r.HandleFunc("/users/{id:[0-9]+}", user.Show).Methods("GET").Name("users.show")

	r.Use(middlewares.StartSession)
}
