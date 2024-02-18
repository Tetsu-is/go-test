package middleware

import (
	"context"
	"net/http"

	"github.com/mileusna/useragent"
)

func UserAgent(h http.Handler) http.Handler {

	type key string

	const (
		osKey key = "OS"
	)

	fn := func(w http.ResponseWriter, r *http.Request) {
		userAgents := r.UserAgent()
		ua := useragent.Parse(userAgents)
		ctx := context.WithValue(r.Context(), osKey, ua.OS)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
