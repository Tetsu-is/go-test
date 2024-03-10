package router

import (
	"api/handler"
	"api/handler/middleware"
	"api/service"
	"database/sql"
	"net/http"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	todoService := service.NewTODOService(todoDB)
	userService := service.NewUserService(todoDB)

	mux := http.NewServeMux()
	mux.Handle("/todos", middleware.ValidToken(middleware.AccessLogger(handler.NewTODOHandler(todoService))))
	mux.Handle("/do-panic", middleware.Recovery(handler.NewPanicHandler()))
	mux.Handle("/hello", handler.NewHelloHandler())
	mux.Handle("/test", middleware.Auth(handler.NewTestHandler()))
	mux.Handle("/auth/signup", handler.NewSignUpHandler(userService))
	mux.Handle("/auth/login", handler.NewLoginHandler(userService))
	mux.Handle("/dev/users", handler.NewGetUserHandler(userService))
	return mux
}
