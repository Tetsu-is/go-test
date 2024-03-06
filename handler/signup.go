package handler

import (
	"api/model"
	"api/service"
	"context"
	"encoding/json"
	"net/http"
)

type SignUpHandler struct {
	svc *service.UserService
}

func NewSignUpHandler(svc *service.UserService) *SignUpHandler {
	return &SignUpHandler{svc: svc}
}

func (h *SignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		req := &model.SignUpRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if req.UserName == "" {
			http.Error(w, "UserName is empty", http.StatusBadRequest)
			return
		}
		if req.Email == "" {
			http.Error(w, "Email is empty", http.StatusBadRequest)
			return
		}
		if req.Password == "" {
			http.Error(w, "Password is empty", http.StatusBadRequest)
			return
		}
		res, err := h.SignUp(r.Context(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *SignUpHandler) SignUp(ctx context.Context, req *model.SignUpRequest) (*model.SignUpResponse, error) {
	usr, err := h.svc.RegisterUser(ctx, req.UserName, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &model.SignUpResponse{
		User: *usr,
	}, nil
}
