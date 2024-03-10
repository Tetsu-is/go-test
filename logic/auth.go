package logic

import (
	"api/model"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

func CreateJwtToken(userID int64) (string, error) {

	//headerとpayloadを作成
	header := model.Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	payload := model.Payload{
		UserID: userID,
		Exp:    time.Now().Add(time.Hour),
	}

	//headerとpayloadをjsonに変換
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	//headerとpayloadをURL safeなbase64に変換
	headerBase64 := base64.RawURLEncoding.EncodeToString(headerJSON)   // 1
	payloadBase64 := base64.RawURLEncoding.EncodeToString(payloadJSON) // 2

	//秘密鍵を定義
	secret := []byte("secret_key")
	signature := hmac.New(sha256.New, secret)
	signature.Write([]byte(strings.Join([]string{headerBase64, payloadBase64}, "."))) // 3 (1と2を.で結合して署名を作成)

	signatureBase64 := base64.RawURLEncoding.EncodeToString(signature.Sum(nil))

	token := strings.Join([]string{headerBase64, payloadBase64, signatureBase64}, ".") // 1, 2, 3を.で結合してトークンを作成

	return token, nil
}

func ResolveJwtToken(token string) (*model.Header, *model.Payload, error) {
	splitToken := strings.Split(token, ".")
	if len(splitToken) != 3 {
		return nil, nil, errors.New("invalid token")
	}

	headerBytes, err := base64.RawURLEncoding.DecodeString(splitToken[0])
	if err != nil {
		return nil, nil, err
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(splitToken[1])
	if err != nil {
		return nil, nil, err
	}

	var header model.Header
	err = json.Unmarshal(headerBytes, &header)
	if err != nil {
		return nil, nil, err
	}
	var payload model.Payload
	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		return nil, nil, err
	}

	return &header, &payload, nil
}
