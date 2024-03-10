package middleware

import (
	"api/handler"
	"api/logic"
	"net/http"
	"time"
)

func CheckToken(h *handler.TODOHandler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, payload, err := logic.ResolveJwtToken(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusInternalServerError)
			return
		}

		userID, exp := payload.UserID, payload.Exp

		//tokenが期限切れならUnAuthorizedを返す
		if exp.Compare(time.Now()) < 0 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		//tokenのuserIDが存在しない場合はUnAuthorizedを返す
		ok := logic.ExistsID(userID, h.svc.CreateUserRepository()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
