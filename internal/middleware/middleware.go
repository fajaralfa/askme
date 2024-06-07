package middleware

import (
	"log"
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
			handler.ApiUnauthorizedErr(w, "anda belum masuk!", nil)
			return
		}
		tokenString = tokenString[len("Bearer "):]
		token, err := jwt.Parse(tokenString)

		if err != nil {
			handler.ApiBadRequestErr(w, "auth token invalid", nil)
			log.Println(err)
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
			return
		}

		handler.ApiUnauthorizedErr(w, "auth token invalid", nil)
	})
}

func PathLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		next.ServeHTTP(w, r)
	})
}
