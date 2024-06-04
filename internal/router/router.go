package router

import (
	"net/http"

	"github.com/fajaralfa/askme/internal/handler"
	"github.com/fajaralfa/askme/internal/helper"
	mw "github.com/fajaralfa/askme/internal/middleware"
	"github.com/fajaralfa/askme/internal/repo"
	"github.com/fajaralfa/askme/internal/service/db"
	"github.com/gorilla/mux"
)

func Create() *mux.Router {
	router := mux.NewRouter()

	// dependency creation
	dbConn := db.Connect()
	userRepo := &repo.User{Db: dbConn}
	qRepo := &repo.Question{Db: dbConn}
	authHr := &handler.Auth{UserRepo: userRepo}
	userHr := &handler.User{UserRepo: userRepo}
	qHr := &handler.Question{QRepo: qRepo, UserRepo: userRepo}

	// api routes
	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(mw.Api)
	api.Methods("POST").Path("/register").HandlerFunc(authHr.Register)
	api.Methods("POST").Path("/login").HandlerFunc(authHr.Login)
	api.Methods("POST").Path("/questions").HandlerFunc(qHr.AskQuestion)
	api.Methods("GET").Path("/questions").Handler(helper.Middlewares(http.HandlerFunc(qHr.FindAllAssociatedWithUser), mw.Authenticated))
	api.Methods("DELETE").Path("/questions/{id:[0-9]+}").Handler(helper.Middlewares(http.HandlerFunc(qHr.RemoveAssociatedWithUser), mw.Authenticated))
	api.Methods("GET").Path("/users/{email:[!-~]+}").HandlerFunc(userHr.FindByEmail)
	api.Methods("GET").Path("/me").HandlerFunc(userHr.CurrentUserInfo)

	// page routes
	router.Use(mw.PathLogger)
	router.PathPrefix("/").HandlerFunc(handler.SpaHandler)

	return router
}
