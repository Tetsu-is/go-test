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
	mux := http.NewServeMux()
	mux.Handle("/todos", middleware.UserAgent(middleware.AccessLogger(handler.NewTODOHandler(todoService))))
	mux.Handle("/do-panic", middleware.Recovery(handler.NewPanicHandler()))
	mux.Handle("/hello", handler.NewHelloHandler())
	mux.Handle("/test", middleware.Auth(handler.NewTestHandler()))
	return mux
}
