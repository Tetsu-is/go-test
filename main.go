package main

import (
	"api/db"
	"api/handler/router"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

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

	server := http.Server{
		Addr:    defaultPort,
		Handler: mux,
	}

	errCh := make(chan error)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done(): // os.Interruptが来た段階で読み取り可能になる -> ctx.Done()でclosedされたチャネルが返される -> shutdownWithTimeout()が実行される??
		shutdownWithTimeout(&server)
	case err := <-errCh:
		log.Printf("server error: %v", err)
	}

	stop()
	//stop()の後はos.Interrupt, os.Kill後の既存の処理が実行される
}

func shutdownWithTimeout(srv *http.Server) {
	//Contextを作成し、30秒のTimeoutを設定
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//Timeoutを設定したContextを渡すことで無期限に待機しないようにする
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

//MEMO
//Server.Shutdown()はActiveな通信を阻害することなく、ListenAndServe()を停止する関数である。

//編集

//devブランチ
