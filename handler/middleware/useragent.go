package middleware

import (
	"context"
	"net/http"

	"github.com/mileusna/useragent"
)

type ctxKey string

const (
	osKey ctxKey = "OS"
)

func SetOS(r *http.Request) *http.Request {
	userAgent := r.UserAgent()
	ua := useragent.Parse(userAgent)
	ctx := context.WithValue(r.Context(), osKey, ua.OS)
	r = r.WithContext(ctx)
	return r
}

func GetOS(r *http.Request) string {
	osValue := r.Context().Value(osKey)
	if osValue == nil {
		return ""
	}
	return r.Context().Value(osKey).(string)
}

func UserAgent(h http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		r = SetOS(r)
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
