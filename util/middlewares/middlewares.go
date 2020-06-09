package middlewares

import (
	"net/http"
	"os"

	"github.com/nazyli/api-restaurant/util/auth"
	"github.com/nazyli/api-restaurant/util/responses"
)

// SetMiddlewareJSON . . .
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errMsg := AppParams(w, r)
		if errMsg != "" {
			responses.ERROR(w, http.StatusNotFound, errMsg)
			return
		}
		next(w, r)
	}
}

// SetMiddlewareAuthentication . . .
func SetMiddlewareAuthentication(next http.HandlerFunc, scopes ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		errMsg := auth.TokenValid(r, scopes...)
		if errMsg != "" {
			responses.ERROR(w, http.StatusUnauthorized, "Unauthorized, "+errMsg)
			return
		}
		errMsg = AppParams(w, r)
		if errMsg != "" {
			responses.ERROR(w, http.StatusNotFound, errMsg)
			return
		}
		next(w, r)
	}
}

func AppParams(w http.ResponseWriter, r *http.Request) (errMsg string) {
	var (
		getParam = r.URL.Query()
		app      = getParam.Get("app_id")
	)

	if app != os.Getenv("APP_NAME") {
		return "404 page not found"
	}
	return ""
}
