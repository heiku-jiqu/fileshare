package main

import (
	"net/http"

	"github.com/heiku-jiqu/fileshare/middleware"
)

func NewFilesRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{user}/files/", unimplemented)
	mux.HandleFunc("POST /{user}/file/", unimplemented)    // initiate new upload
	mux.HandleFunc("PUT /{user}/file/{id}", unimplemented) // complete upload?
	return middleware.Logger(mux)
}
