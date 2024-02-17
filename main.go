package main

import (
	"net/http"
	"test/db"
	"test/handler/router"
)

func main() {
	const (
		defaultPort   = ":8080"
		defaultDBPath = ".sqlite3/todo.db"
	)

	todoDB, err := db.NewDB(defaultDBPath)
	if err != nil {
		panic(err)
	}
	defer todoDB.Close()

	mux := router.NewRouter(todoDB)

	http.ListenAndServe(defaultPort, mux)

}
