package handler

import (
	"api/model"
	"api/service"
	"context"
	"encoding/json"
	"net/http"
)

type LoginHandler struct {
	svc *service.UserService
}

func NewLoginHandler(svc *service.UserService) *LoginHandler {
	return &LoginHandler{svc: svc}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		req := &model.LogInRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if req.Email == "" || req.Password == "" {
			http.Error(w, "Email or Password is Empty", http.StatusBadRequest)
			return
		}
		res, err := h.Login(r.Context(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cookie := &http.Cookie{
			Name:  "token",
			Value: res.Token,
		}

		http.SetCookie(w, cookie)

		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *LoginHandler) Login(ctx context.Context, req *model.LogInRequest) (*model.LogInResponse, error) {
	token, err := h.svc.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &model.LogInResponse{
		Token: token,
	}, nil
}
