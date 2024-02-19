package middleware

import (
	"net/http"
	"os"
)

func checkAuth(r *http.Request) bool {

	userId, password, ok := r.BasicAuth()
	if !ok {
		return false
	}
	return userId == os.Getenv("BASIC_AUTH_USER_ID") && password == os.Getenv("BASIC_AUTH_PASSWORD")
}

func Auth(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !checkAuth(r) {
			w.Header().Set("WWW-Authenticate", `Basic realm="WallyWorld"`)
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
