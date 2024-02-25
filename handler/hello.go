package handler

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HelloHandler is called!!"))
	fmt.Println("HelloHandler is called")
}
