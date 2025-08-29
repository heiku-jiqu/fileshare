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

func unimplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "unimplemented")
}
