package handlers

import (
	"net/http"
	"time"

	"github.com/bigmassa/gohttp-starter/internal/conf"
	"github.com/bigmassa/gohttp-starter/internal/models"

	"github.com/go-playground/validator/v10"
)

type loginForm struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Errors   map[string]string
}

func (f *loginForm) Validate() bool {
	f.Errors = make(map[string]string)
	msgs := map[string]string{
		"required": "This field is required.",
		"email":    "Please enter a valid email address.",
	}
	if err := conf.GetValidate().Struct(f); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			f.Errors[err.Field()] = msgs[err.Tag()]
		}
	}
	return len(f.Errors) == 0
}

func LoginGetHandler(a *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	context := struct {
		Form *loginForm
	}{
		Form: &loginForm{},
	}
	return a.TemplateResponse(w, "loginHTML", context)
}

func LoginPostHandler(a *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	form := &loginForm{
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if form.Validate() == true {
		user := &models.User{}
		if err := a.Db.Where("email = ?", form.Email).First(&user).Error; err == nil {
			// check the users password and set the session and redirect if correct
			if user.CanLogin() && user.CheckPassword(form.Password) {
				// sets the session cookie
				session, _ := a.CookieStore.Get(r, "session")
				session.Values["user_id"] = user.ID
				session.Save(r, w)
				// update the users last logged in at time
				user.LoggedInAt = time.Now()
				a.Db.Save(&user)
				// finally redirect them
				http.Redirect(w, r, "/", http.StatusFound)
				return http.StatusFound, nil
			}
		}
	}
	// otherwise render the form with errors and provide a default error for password
	// if it doesnt exist
	if _, ok := form.Errors["Password"]; ok == false {
		form.Errors["Password"] = "Please check your details and try again"
	}
	context := struct {
		Form *loginForm
	}{
		Form: form,
	}
	return a.TemplateResponse(w, "loginHTML", context)
}
