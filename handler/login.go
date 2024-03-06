package handler

import (
	"api/service"
	"net/http"
)

type LoginHandler struct {
	svc *service.UserService
}

func NewLoginHandler(svc *service.UserService) *LoginHandler {
	return &LoginHandler{svc: svc}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
