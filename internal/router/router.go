package router

import (
	"github.com/fajaralfa/askme/internal/feature/spa"
	"github.com/fajaralfa/askme/internal/feature/user"
	"github.com/fajaralfa/askme/internal/middleware"
	"github.com/gorilla/mux"
)

func Create() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.Api)
	api.HandleFunc("/users", user.GetHandler)
	api.HandleFunc("/users/{id:[0-9]+}", user.FindHandler)

	router.PathPrefix("/").HandlerFunc(spa.SpaHandler)

	return router
}
