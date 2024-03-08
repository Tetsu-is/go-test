package logic

import (
	"api/model"
	"encoding/base64"
	"encoding/json"
	"time"
)

func createJwtToken(userID int64) (model.Jwt, error) {
	header := model.Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	payload := model.Payload{
		UserID: userID,
		Exp:    time.Now().Add(time.Hour * 24),
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return model.Jwt{}, err
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return model.Jwt{}, err
	}

	headerBase64 := base64.StdEncoding.EncodeToString(headerJSON)
	payloadBase64 := base64.StdEncoding.EncodeToString(payloadJSON)

	 := base64.StdEncoding.EncodeToString([]byte(haederBase64 + "." + payloadBase64))
	return model.Jwt{}, nil
}
