package middleware

import (
	"log/slog"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip     = r.RemoteAddr
			method = r.Method
			url    = r.URL.String()
			proto  = r.Proto
		)
		next.ServeHTTP(w, r)
		slog.Info("Request served.", "protocol", proto, "method", method, "url", url, "ip", ip, "matched", r.Pattern)
	})

}
