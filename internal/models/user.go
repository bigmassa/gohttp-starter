package models

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bigmassa/gohttp-starter/internal/utils"

	"github.com/binaryknights/gutil"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email      string `gorm:"type:varchar(120);unique_index;not null"`
	FirstName  string `gorm:"type:varchar(150);not null"`
	LastName   string `gorm:"type:varchar(150);not null"`
	Password   string `gorm:"type:varchar(255);not null"`
	IsActive   bool   `gorm:"type:boolean;not null;default:true"`
	LoggedInAt time.Time
}

func (u *User) CanLogin() bool {
	return u.IsActive
}

func (u *User) CheckPassword(password string) bool {
	return gutil.ComparePasswords(u.Password, []byte(password))
}

func (u *User) SendPasswordReset(r *http.Request) error {
	token, _ := utils.NewPasswordResetToken(u.Password, 12*time.Hour)
	from := os.Getenv("EMAIL_DEFAULT_FROM_ADDRESS")
	to := []string{u.Email}
	subject := "Password reset at " + r.Host
	body := ""
	email := gutil.NewMail(from, to, subject, body)
	data := struct {
		Token string
		Host  string
	}{
		Token: token,
		Host:  r.Host,
	}
	if err := email.ParseTemplate("web/templates/email/password_reset.txt", data); err == nil {
		if err := email.SendMail(); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (u *User) SetPassword(password string) {
	hash := gutil.CreatePassword([]byte(password))
	u.Password = string(hash)
}

func (u *User) String() string {
	return u.Email
}
