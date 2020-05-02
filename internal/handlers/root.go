package handlers

import (
	"net/http"

	"github.com/bigmassa/gohttp-starter/internal/conf"
	"github.com/bigmassa/gohttp-starter/internal/middleware"
	"github.com/bigmassa/gohttp-starter/internal/models"
)

func IndexHandler(a *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	context := struct {
		User *models.User
	}{
		User: middleware.GetUser(r.Context()),
	}
	return a.TemplateResponse(w, "indexHTML", context)
}
