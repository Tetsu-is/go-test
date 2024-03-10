package middleware

import (
	"api/logic"
	"context"
	"net/http"
	"time"
)

func ValidToken(h http.Handler) http.Handler {
	//tokenが正しいか
	//認可はserviceで行う

	type ctxKey string
	const tokenKey ctxKey = "token" //built-inの型を避けるため

	fn := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		header, payload, err := logic.ResolveJwtToken(cookie.Value)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusInternalServerError)
			return
		}

		if header.Alg != "HS256" || header.Typ != "JWT" {
			http.Error(w, "header is invalid", http.StatusUnauthorized)
			return
		}

		if payload.Exp.Compare(time.Now()) < 0 {
			http.Error(w, "token is expired", http.StatusUnauthorized)
			return
		}

		//tokenをcontextに格納
		ctx := context.WithValue(r.Context(), tokenKey, cookie.Value)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
