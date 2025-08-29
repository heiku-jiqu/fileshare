package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	t "github.com/heiku-jiqu/fileshare/web/template"
)

func main() {
	log.Print("Hello World!")

	mux := http.NewServeMux()
	mux.Handle("GET /", http.FileServerFS(t.Website))
	mux.HandleFunc("GET /api/v1/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
