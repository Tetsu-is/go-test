package handler

import (
	"api/model"
	"api/service"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

		q := r.URL.Query()
		offsetID := q.Get("offset_id")
		limit := q.Get("limit")

		if offsetID == "" {
			fmt.Println("offsetID is empty but it's okay")
		}
		if limit == "" {
			fmt.Println("limit is empty but it's okay")
		}

		parsedOffsetID, _ := strconv.ParseInt(offsetID, 10, 64) //default value is 0
		parsedLimit, _ := strconv.ParseInt(limit, 10, 64)       //default value is 0

		req.OffsetID, req.Limit = parsedOffsetID, parsedLimit //Set offsetID and limit from QueryParam

		res, err := h.Read(r.Context(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err, "err handler read")
			return
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err, "err handler encode")
			return
		}
	}
}

func (h *GetUserHandler) Read(ctx context.Context, req *model.GetUserRequest) (*model.GetUserResponse, error) {
	users, err := h.svc.ReadUser(ctx, req.OffsetID, req.Limit)
	if err != nil {
		return nil, err
	}
	return &model.GetUserResponse{User: users}, nil
}
