package main

import (
	"fmt"
	"net/http"
	"strconv"

	f "github.com/heiku-jiqu/fileshare/appsvc/files"
	"github.com/heiku-jiqu/fileshare/model/file"
	"github.com/heiku-jiqu/fileshare/model/user"
)

type handler struct {
	app *f.FilesApp
}

func NewFilesRouter() http.Handler {
	mux := http.NewServeMux()
	app := f.NewFilesApp(f.NewFilesDB())
	h := &handler{app: app}
	mux.HandleFunc("GET /{user}/files", h.GetFiles)
	mux.HandleFunc("POST /{user}/file", unimplemented)     // initiate new upload
	mux.HandleFunc("PUT /{user}/file/{id}", unimplemented) // complete upload?
	mux.HandleFunc("/healthcheck", Healthcheck)
	return (mux)
}

func (h *handler) GetFiles(w http.ResponseWriter, r *http.Request) {
	files, err := h.app.GetLatest(r.Context(), 5)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for i, file := range files {
		fmt.Fprintf(w, "%d: %s\t%s\n", i, file.Name, file.Status)
	}
}

func (h *handler) PostFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	name := r.PostForm.Get("name")
	sizeString := r.PostForm.Get("size")
	size, err := strconv.Atoi(sizeString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	new := file.NewFile(name, int64(size), user.UserId(2))
	h.app.Insert(r.Context(), new)
}
