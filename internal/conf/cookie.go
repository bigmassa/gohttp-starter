package conf

import (
	"sync"

	"github.com/binaryknights/gutil"
	"github.com/gorilla/sessions"
)

var cookieStore *sessions.CookieStore
var cookieStoreOnce sync.Once

func GetCookieStore() *sessions.CookieStore {
	cookieStoreOnce.Do(func() {
		cookieStore = sessions.NewCookieStore(
			gutil.DecodeEnvKey("SESSION_AUTH_KEY"),
			gutil.DecodeEnvKey("SESSION_ENCRYPTION_KEY"),
		)
	})
	return cookieStore
}
