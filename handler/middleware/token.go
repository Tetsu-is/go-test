package middleware

import "net/http"

func checkToken(h http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
