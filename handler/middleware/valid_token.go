package middleware

import (
	"api/handler"
	"api/logic"
	"net/http"
	"time"
)

func ValidToken(h *handler.TODOHandler) http.Handler {
	//tokenが正しいか
	//認可はserviceで行う

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

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
