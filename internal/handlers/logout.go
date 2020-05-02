package handlers

import (
	"net/http"

	"github.com/bigmassa/gohttp-starter/internal/conf"
)

func LogoutHandler(a *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	session, _ := a.CookieStore.Get(r, "session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
	return http.StatusFound, nil
}
