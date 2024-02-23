package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/model"
	"time"
)

func AccessLogger(h http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		h.ServeHTTP(w, r)

		end := time.Now()
		latency := end.Sub(start).Milliseconds()
		os := GetOS(r)
		accesslog := &model.AccessLog{
			Timestamp: start,
			Latency:   latency,
			Path:      r.URL.Path,
			OS:        os,
		}
		data, err := json.Marshal(accesslog)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data))
	}
	return http.HandlerFunc(fn)
}
