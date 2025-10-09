package main

import (
	"fmt"
	"net/http"

	f "github.com/heiku-jiqu/fileshare/appsvc/files"
)

func NewFilesRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{user}/files/", GetFiles)
	mux.HandleFunc("POST /{user}/file/", unimplemented)    // initiate new upload
	mux.HandleFunc("PUT /{user}/file/{id}", unimplemented) // complete upload?
	return (mux)
}

func GetFiles(w http.ResponseWriter, r *http.Request) {
	files := f.ListFiles()
	for i, file := range files {
		fmt.Fprintf(w, "%d: %s\t%s\n", i, file.Name, file.Status)
	}
}
