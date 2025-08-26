package file

import (
	"testing"

	"github.com/heiku-jiqu/fileshare/user"
)

func TestFile(t *testing.T) {
	filename := "example.txt"
	const filesize int64 = 20 * 1024 * 1024 // 20 MiB
	userId := user.UserId(1)
	file := NewFile(filename, filesize, userId)
	if file.name != filename {
		t.Errorf("Expected filename %s, got %s", filename, file.name)
	}
	if Started != file.status {
		t.Errorf("Expected status %s, got %s", Started, file.status)
	}

	uploadInfo, err := blobStore.GeneratePresignedUploadURLs(file)

	// pendingParts := file.PendingParts()
	// file.isPartUploaded(parts[0])

	// fileStatus := file.MarkPartUploaded(partNum, etag) // mark part as uploaded when s3 part uploaded event is received. returns file state
	// internally checks whether all parts have been uplaoded

	// status := file.Status()
	// file.MarkCompleted(storageURL)

	url, err := blobStore.GenerateDownloadURL(file) // err if file has not completed upload

	// errors and retries:
	// what happens when aborted?
	// what happens when part fails?
	// what happens when server restarts?
	// what happens when upload URL expires?

}
