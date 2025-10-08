package main

import (
	"github.com/heiku-jiqu/fileshare/file"
	"github.com/heiku-jiqu/fileshare/user"
)

func filesDB() []file.File {
	return []file.File{
		file.NewFile("abc", 123, user.UserId(1)),
		file.NewFile("def", 456, user.UserId(1)),
		file.NewFile("xyz", 789, user.UserId(2)),
	}

}

func NewFileApp() {

}

func ListFiles() []file.File {
	return []file.File{
		file.NewFile("abc", 123, user.UserId(1)),
		file.NewFile("def", 456, user.UserId(1)),
		file.NewFile("abc", 123, user.UserId(1)),
	}
}
