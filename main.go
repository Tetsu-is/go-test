package main

import (
	"log"
	"net/http"
	"test/db"
	"test/handler/router"

	"github.com/joho/godotenv"
)

func main() {
	const (
		defaultPort   = ":8080"
		defaultDBPath = ".sqlite3/todo.db"
	)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	todoDB, err := db.NewDB(defaultDBPath)
	if err != nil {
		panic(err)
	}
	defer todoDB.Close()

	mux := router.NewRouter(todoDB)

	http.ListenAndServe(defaultPort, mux)

}
