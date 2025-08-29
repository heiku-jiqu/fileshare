package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/heiku-jiqu/fileshare/middleware"
	t "github.com/heiku-jiqu/fileshare/web/template"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServerFS(t.Website))
	mux.HandleFunc("/login", unimplemented)
	mux.HandleFunc("/healthcheck", Healthcheck)
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", NewFilesRouter()))

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

func NewFilesRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{user}/files/", unimplemented)
	mux.HandleFunc("POST /{user}/file/", unimplemented)    // initiate new upload
	mux.HandleFunc("PUT /{user}/file/{id}", unimplemented) // complete upload?
	return middleware.Logger(mux)
}

func unimplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "unimplemented")
}
