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
	mux.Handle("/todos", handler.NewTODOHandler(todoService))
	// mux.HandleFunc("/do-panic", handler.NewPanicHandler().ServeHTTP)
	mux.Handle("/do-panic", middleware.Recovery(handler.NewPanicHandler()))
	return mux
}
