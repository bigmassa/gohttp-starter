package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bigmassa/gohttp-starter/internal/conf"
	"github.com/bigmassa/gohttp-starter/internal/models"
	"github.com/bigmassa/gohttp-starter/internal/routing"

	"github.com/gorilla/mux"
)

func init() {
	// wait for a db connection before we continue
	// useful when waiting for a docker container to become ready
	conf.WaitForDB()
}

func main() {
	listen := flag.String("listen", ":80", "listen addr eg :80")

	flag.Parse()

	// setup database
	db, _ := conf.ConnectDB()
	defer db.Close()
	models.SetDatabase(db)

	// setup global app context
	context := &conf.AppContext{
		Db:          db,
		CookieStore: conf.GetCookieStore(),
		Templates:   conf.GetTemplates(),
		Validator:   conf.GetValidate(),
	}

	// routing
	r := mux.NewRouter()
	// routes
	r = routing.BaseRoutes(r, context)
	r = routing.StaticRoutes(r, context)
	r = routing.AuthRoutes(r, context)
	r = routing.SampleLockedRoutes(r, context)

	http.Handle("/", r)

	// serve
	log.Fatal(http.ListenAndServe(*listen, nil))
}
