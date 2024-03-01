package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"test/db"
	"test/handler/router"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	ctxWaitSignal, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

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

	server := http.Server{
		Addr:    defaultPort,
		Handler: mux,
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		server.ListenAndServe()
	}()

	//signalを待つctxWaitSignalを終了させる
	<-ctxWaitSignal.Done()

	//シャットダウン用のコンテキストを作成し、30秒のタイムアウトを設定
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctxShutDown); err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
