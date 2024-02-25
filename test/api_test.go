package test

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// Basic Auth Test
// 正常系
// 対象の API のみ Basic 認証がかかっているか、どうか。
// 正しい User ID, Password で Basic 認証をクリアしアクセスできるかどうか。

// 異常系
// 間違った User ID, Password を送信した場合、 Basic 認証が失敗しHTTP Status Code が 401 で返却されているかどうか。
// 空の User ID, Password を送信した場合、 Basic 認証が失敗し HTTP Status Code が 401 で返却されているかどうか。
// アクセス時に User ID, Password を送信しなかった場合、Basic 認証が失敗し HTTP Status Code が 401 で返却されているかどうか。

func Test1(t *testing.T) {
	t.Parallel()

	if err := godotenv.Load("../.env"); err != nil {
		t.Errorf("Error loading .env file")
	}

	const baseUrl = "http://localhost:8080"

	testcases := map[string]struct {
		Endpoints    string
		RequestType  string
		UserID       string
		Password     string
		ExpectStatus int
	}{

		"No_Auth_for /todos": {
			Endpoints:    "/hello",
			RequestType:  "GET",
			ExpectStatus: 200,
		},

		"No_Auth_for /do-panic": {
			Endpoints:    "/do-panic",
			RequestType:  "GET",
			ExpectStatus: 200,
		},

		"No_Auth_for /hello": {
			Endpoints:    "/hello",
			RequestType:  "GET",
			ExpectStatus: 200,
		},

		"Need_Auth_for /test": {
			Endpoints:    "/test",
			RequestType:  "GET",
			ExpectStatus: 401,
		},

		"Auth Success (CorrectUserID_CorrectPassword)": {
			Endpoints:    "/test",
			RequestType:  "GET",
			UserID:       os.Getenv("BASIC_AUTH_USER_ID"),
			Password:     os.Getenv("BASIC_AUTH_PASSWORD"),
			ExpectStatus: 200,
		},

		"Auth Fail (WrongUserID_CorrectPassword)": {
			Endpoints:    "/test",
			RequestType:  "GET",
			UserID:       "wrongUserID",
			Password:     os.Getenv("BASIC_AUTH_PASSWORD"),
			ExpectStatus: 401,
		},

		"Auth Fail (CorrectUserID_WrongPassword)": {
			Endpoints:    "/test",
			RequestType:  "GET",
			UserID:       os.Getenv("BASIC_AUTH_USER_ID"),
			Password:     "wrongPassword",
			ExpectStatus: 401,
		},

		"Auth Fail (WrongUserID_WrongPassword)": {
			Endpoints:    "/test",
			RequestType:  "GET",
			UserID:       "wrongUserID",
			Password:     "wrongPassword",
			ExpectStatus: 401,
		},

		"Auth Fail (EmptyUserID_EmptyPassword)": {
			Endpoints:    "/test",
			RequestType:  "GET",
			UserID:       "",
			Password:     "",
			ExpectStatus: 401,
		},
	}

	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			url := baseUrl + tc.Endpoints
			//base64 conversion of userID and password
			encoded := base64.StdEncoding.EncodeToString([]byte(tc.UserID + ":" + tc.Password))

			req, err := http.NewRequest(tc.RequestType, url, nil)
			if err != nil {
				t.Errorf("Error creating request: %v", err)
			}

			req.Header.Set("Authorization", "Basic "+encoded)

			client := new(http.Client)
			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Error sending request: %v", err)
			}

			if resp.StatusCode != tc.ExpectStatus {
				fmt.Println("tried ID: ", tc.UserID, "tried Password: ", tc.Password, "tried URL: ", url, "tried RequestType: ", tc.RequestType, "tried ExpectStatus: ", tc.ExpectStatus, "got status: ", resp.StatusCode)
				t.Errorf("Expected status code %d, but got %d", tc.ExpectStatus, resp.StatusCode)
			}

			defer resp.Body.Close()
		})
	}

}
