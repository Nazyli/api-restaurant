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
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Add("Access-Control-Expose-Headers", "responseType")
		w.Header().Add("Access-Control-Expose-Headers", "observe")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, PATCH,OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
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
