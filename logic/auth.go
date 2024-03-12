package logic

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var mySigningKey = []byte("secret_key")

func CreateJwtToken(userID int64) (string, error) {

	claims := jwt.MapClaims{
		"user_id": strconv.FormatInt(userID, 10),
		"exp":     time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func VerifyJwtToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return jwt.ErrSignatureInvalid
	}

	return nil
}
