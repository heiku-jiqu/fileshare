package files

import (
	"context"

	"github.com/heiku-jiqu/fileshare/model/file"
	"github.com/heiku-jiqu/fileshare/model/user"
)

// FilesRepo provides a Repository to store / get File.
type FilesRepo interface {
	// GetLatest list of uploaded files.
	GetLatest(ctx context.Context, num int) ([]file.File, error)
	// Insert a new file
	Insert(ctx context.Context, f file.File) error
}

// FilesApp is an application to upload and retrieve files.
type FilesApp struct {
	// Embed FilesRepo for now because still CRUD.
	// When there is a need to add an application layer indirection, refactor the embed.
	FilesRepo
}

func NewFilesApp(repo FilesRepo) *FilesApp {
	app := &FilesApp{FilesRepo: repo}
	app.Insert(context.TODO(), file.NewFile("abc", 123, user.UserId(1)))
	app.Insert(context.TODO(), file.NewFile("def", 456, user.UserId(1)))
	app.Insert(context.TODO(), file.NewFile("xyz", 789, user.UserId(2)))
	return app
}
