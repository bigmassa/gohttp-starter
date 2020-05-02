package utils

import (
	"time"

	"github.com/binaryknights/gutil"
	"github.com/dchest/passwordreset"
)

func GetTokenHash(login string) ([]byte, error) {
	return []byte(login), nil
}

func NewPasswordResetToken(login string, dur time.Duration) (token string, err error) {
	hash, err := GetTokenHash(login)
	if err != nil {
		return
	}
	secret := gutil.DecodeEnvKey("PASSWORD_RESET_KEY")
	token = passwordreset.NewToken(login, dur, hash, secret)
	return
}

func VerifyPasswordResetToken(token string) (login string, err error) {
	secret := gutil.DecodeEnvKey("PASSWORD_RESET_KEY")
	login, err = passwordreset.VerifyToken(token, GetTokenHash, secret)
	return
}
