package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/heiku-jiqu/fileshare/middleware"
	"github.com/heiku-jiqu/fileshare/web"
)

func main() {
	sessionManager := scs.New()
	login := NewLogin(sessionManager)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.FileServerFS(web.Static))
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		web.Index.Execute(w, nil)
	})
	mux.HandleFunc("/healthcheck", Healthcheck)
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", NewFilesRouter()))
	mux.HandleFunc("POST /login", login.LoginPostHandler)
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		web.Login.Execute(w, nil)
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        middleware.Logger(sessionManager.LoadAndSave(mux)),
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
