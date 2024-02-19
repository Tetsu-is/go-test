package handler

import (
	"fmt"
	"net/http"
)

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TestHandler is called"))
	fmt.Println("TestHandler is called")
}
