package files

import (
	"context"

	"github.com/heiku-jiqu/fileshare/model/file"
)

// FilesDB implements the interface FilesInterface that is required by FilesApp
type FilesDB struct {
	store  map[int]file.File
	currId int
}

func NewFilesDB() *FilesDB {
	store := make(map[int]file.File)
	return &FilesDB{store: store, currId: 0}
}

func (db *FilesDB) Insert(ctx context.Context, f file.File) error {
	db.store[db.currId] = f
	db.currId += 1
	return nil
}

func (db *FilesDB) GetLatest(ctx context.Context, num int) ([]file.File, error) {
	out := make([]file.File, num)
	for i := range num {
		idx := db.currId - i
		out[i] = db.store[idx]
	}
	return out, nil
}
