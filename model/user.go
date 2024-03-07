package model

import "time"

type (
	User struct {
		ID        int64     `json:"id"`
		UserName  string    `json:"user_name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Token struct {
		ID        int64     `json:"id"`
		UserID    int64     `json:"user_id"`
		Token     string    `json:"token"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	SignUpRequest struct {
		UserName string `json:"user_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignUpResponse struct {
		User User `json:"user"`
	}

	LogInRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LogInResponse struct {
		Token string `json:"token"`
	}
)
