package router

import (
	"database/sql"
	"net/http"
	"test/handler"
	"test/handler/middleware"
	"test/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	todoService := service.NewTODOService(todoDB)
	mux := http.NewServeMux()
	mux.Handle("/todos", middleware.UserAgent(middleware.AccessLogger(handler.NewTODOHandler(todoService))))
	mux.Handle("/do-panic", middleware.Recovery(handler.NewPanicHandler()))
	mux.Handle("/test", middleware.Auth(handler.NewTestHandler()))
	return mux
}
