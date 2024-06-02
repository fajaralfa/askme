package middleware

import (
	"net/http"

	"github.com/fajaralfa/askme/internal/handler"
	"github.com/fajaralfa/askme/internal/service/jwt"
)

func Api(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			handler.ApiUnauthorizedErr(w, "missing authorization header", nil)
			return
		}
		tokenString = tokenString[len("Bearer "):]
		token, _ := jwt.Parse(tokenString)

		if token.Valid {
			next.ServeHTTP(w, r)
			return
		}

		handler.ApiUnauthorizedErr(w, "auth token invalid", nil)
	})
}
