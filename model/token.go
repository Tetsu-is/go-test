package model

import "time"

type (
	Token struct {
		ID        int64     `json:"id"`
		UserID    int64     `json:"user_id"`
		Token     string    `json:"token"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Header struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}

	Payload struct {
		UserID int64     `json:"user_id"`
		Exp    time.Time `json:"exp"`
	}

	Jwt struct {
		Header struct {
			Alg string `json:"alg"`
			Typ string `json:"typ"`
		}
		Payload struct {
			UserID int64     `json:"user_id"`
			Exp    time.Time `json:"exp"`
		}
		Signature string `json:"signature"`
	}
)
