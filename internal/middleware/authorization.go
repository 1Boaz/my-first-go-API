package middleware

import (
	"errors"
	"net/http"

	api "github.com/1Boaz/my-first-go-API/api"
	"github.com/1Boaz/my-first-go-API/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorized = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorized)
			api.RequestErrorHandler(w, UnAuthorized)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil{
			api.InternalErrorHandler(w)
		}

		var loginDetalis *tools.LoginDetalis
		loginDetalis = (*database).GetUserLoginDetalis(username)

		if loginDetalis == nil || (token != (*loginDetalis).AuthToken) {
			log.Error(UnAuthorized)
			api.RequestErrorHandler(w, UnAuthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}