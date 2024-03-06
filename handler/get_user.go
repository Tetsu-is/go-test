package handler

import (
	"api/model"
	"api/service"
	"net/http"
)

type GetUserHandler struct {
	svc *service.UserService
}

func NewGetUserHandler(svc *service.UserService) *GetUserHandler {
	return &GetUserHandler{svc: svc}
}

func (h *GetUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		req := &model.GetUserRequest{}
	}
}
