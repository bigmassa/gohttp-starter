package routing

import (
	"net/http"

	"github.com/bigmassa/gohttp-starter/internal/conf"
	"github.com/bigmassa/gohttp-starter/internal/handlers"
	"github.com/bigmassa/gohttp-starter/internal/middleware"

	"github.com/gorilla/mux"
)

func BaseRoutes(r *mux.Router, ctx *conf.AppContext) *mux.Router {
	r.Handle("/", middleware.CurrentUser(conf.AppHandler{ctx, handlers.IndexHandler})).Methods("GET")
	return r
}

func AuthRoutes(r *mux.Router, ctx *conf.AppContext) *mux.Router {
	s := r.PathPrefix("/auth").Subrouter()
	s.Handle("/login", conf.AppHandler{ctx, handlers.LoginGetHandler}).Methods("GET")
	s.Handle("/login", conf.AppHandler{ctx, handlers.LoginPostHandler}).Methods("POST")
	s.Handle("/logout", conf.AppHandler{ctx, handlers.LogoutHandler}).Methods("GET")
	s.Handle("/password/reset", conf.AppHandler{ctx, handlers.PWResetGetHandler}).Methods("GET")
	s.Handle("/password/reset", conf.AppHandler{ctx, handlers.PWResetPostHandler}).Methods("POST")
	s.Handle("/password/reset/done", conf.AppHandler{ctx, handlers.PWResetDoneHandler}).Methods("GET")
	s.Handle("/password/reset/confirm", conf.AppHandler{ctx, handlers.PWResetConfirmGetHandler}).Methods("GET")
	s.Handle("/password/reset/confirm", conf.AppHandler{ctx, handlers.PWResetConfirmPostHandler}).Methods("POST")
	s.Handle("/password/reset/complete", conf.AppHandler{ctx, handlers.PWResetCompleteHandler}).Methods("GET")
	return r
}

func StaticRoutes(r *mux.Router, ctx *conf.AppContext) *mux.Router {
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	return r
}

func SampleLockedRoutes(r *mux.Router, ctx *conf.AppContext) *mux.Router {
	s := r.PathPrefix("/").Subrouter()
	s.Use(middleware.CurrentUser, middleware.AuthRequired)
	s.Handle("/restricted", conf.AppHandler{ctx, handlers.IndexHandler}).Methods("GET")
	return r
}
