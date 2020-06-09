package middlewares

import (
	"net/http"

	"github.com/nazyli/api-restaurant/util/auth"
	"github.com/nazyli/api-restaurant/util/responses"
)

// SetMiddlewareJSON . . .
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
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
		next(w, r)
	}
}
