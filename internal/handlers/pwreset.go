package handlers

import (
	"net/http"

	"github.com/bigmassa/gohttp-starter/internal/conf"
	"github.com/bigmassa/gohttp-starter/internal/models"
	"github.com/bigmassa/gohttp-starter/internal/utils"

	"github.com/go-playground/validator/v10"
)

type pwResetForm struct {
	Email  string `validate:"required,email"`
	Errors map[string]string
}

func (f *pwResetForm) Validate() bool {
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

func PWResetGetHandler(a *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	context := struct {
		Form *pwResetForm
	}{
		Form: &pwResetForm{},
	}
	return a.TemplateResponse(w, "pwresetHTML", context)
}

func PWResetPostHandler(a *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	form := &pwResetForm{
		Email: r.PostFormValue("email"),
	}
	if form.Validate() == false {
		context := struct {
			Form *pwResetForm
		}{
			Form: form,
		}
		return a.TemplateResponse(w, "pwresetHTML", context)
	}
	// get the user - for security reasons we dont want to alert the user
	// if the email doesnt exist
	user := &models.User{}
	if err := a.Db.Where("email = ?", form.Email).First(&user).Error; err == nil {
		go func() { user.SendPasswordReset(r) }()
	}
	// finally redirect them even if anything failed
	http.Redirect(w, r, "/auth/password/reset/done", http.StatusFound)
	return http.StatusFound, nil
}

func PWResetDoneHandler(a *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	return a.TemplateResponse(w, "pwresetdoneHTML", nil)
}

type pwResetConfirmForm struct {
	NewPassword    string `validate:"required"`
	RepeatPassword string `validate:"required,eqfield=NewPassword"`
	Errors         map[string]string
}

func (f *pwResetConfirmForm) Validate() bool {
	f.Errors = make(map[string]string)
	msgs := map[string]string{
		"required": "This field is required.",
		"eqfield":  "The passwords do not match.",
	}
	if err := conf.GetValidate().Struct(f); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			f.Errors[err.Field()] = msgs[err.Tag()]
		}
	}
	return len(f.Errors) == 0
}

func getUserFromToken(a *conf.AppContext, r *http.Request) (*models.User, error) {
	token := r.FormValue("token")
	// check the token is valid
	pwhash, err := utils.VerifyPasswordResetToken(token)
	if err != nil {
		return nil, err
	}
	// check the users password is still as it was when they requested a reset
	user := &models.User{}
	if err := a.Db.Where("password = ?", pwhash).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func PWResetConfirmGetHandler(a *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	if _, err := getUserFromToken(a, r); err != nil {
		return http.StatusNotFound, err
	}
	context := struct {
		Form *pwResetConfirmForm
	}{
		Form: &pwResetConfirmForm{},
	}
	return a.TemplateResponse(w, "pwresetconfirmHTML", context)
}

func PWResetConfirmPostHandler(a *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	user, err := getUserFromToken(a, r)
	if err != nil {
		return http.StatusNotFound, err
	}
	// validate the form
	form := &pwResetConfirmForm{
		NewPassword:    r.PostFormValue("new_password"),
		RepeatPassword: r.PostFormValue("repeat_password"),
	}
	if form.Validate() == false {
		context := struct {
			Form *pwResetConfirmForm
		}{
			Form: form,
		}
		return a.TemplateResponse(w, "pwresetconfirmHTML", context)
	}
	user.SetPassword(form.NewPassword)
	if err := models.Db.Save(&user).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	http.Redirect(w, r, "/auth/password/reset/complete", http.StatusFound)
	return http.StatusFound, nil
}

func PWResetCompleteHandler(a *conf.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	return a.TemplateResponse(w, "pwresetcompleteHTML", nil)
}
