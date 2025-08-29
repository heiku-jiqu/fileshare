package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	t "github.com/heiku-jiqu/fileshare/web/template"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServerFS(t.Website))
	mux.HandleFunc("/login", unimplemented)
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", NewApiRouter()))

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	slog.Info("Listening and serving...")
	log.Fatal(s.ListenAndServe())
}

func NewApiRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", healthcheck)
	mux.HandleFunc("GET /{user}/files/", unimplemented)
	mux.HandleFunc("POST /{user}/file/", unimplemented)    // initiate new upload
	mux.HandleFunc("PUT /{user}/file/{id}", unimplemented) // complete upload?
	return logger(mux)
}
func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func logger(next http.Handler) http.Handler {
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

func unimplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "unimplemented")
}
