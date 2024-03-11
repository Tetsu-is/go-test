package middleware

import (
	"api/logic"
	"context"
	"net/http"
)

func ValidToken(h http.Handler) http.Handler {
	//tokenが正しいか
	//認可はserviceで行う

	type ctxKey string
	const tokenKey ctxKey = "token" //built-inの型を避けるため

	fn := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "unauthorized!", http.StatusUnauthorized)
			return
		}

		if err := logic.VerifyJwtToken(cookie.Value); err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		//tokenをcontextに格納
		ctx := context.WithValue(r.Context(), tokenKey, cookie.Value)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
