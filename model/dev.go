package model

type (
	GetUserRequest struct {
		OffsetID int64 `json:"offset_id"`
		Limit    int64 `json:"limit"`
	}
	GetUserResponse struct {
		User []*User
	}
)
