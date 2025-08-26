package file

import "testing"

func TestFile(t *testing.T) {
	file := NewFile(filename, size)
	// urls := file.GenerateUploadURLs()  // browser uses URLs to upload directly to storage
	// file.GenerateRemainingUploadURLs() // urls to upload parts

	// browser needs to know:
	// how to chunk the parts
	// where each chunk should be uploaded
	// which parts have not been uploaded

	pendingParts := file.PendingParts()
	// file.isPartUploaded(parts[0])

	fileStatus := file.MarkPartUploaded(partNum, etag) // mark part as uploaded when s3 part uploaded event is received. returns file state
	// internally checks whether all parts have been uplaoded

	status := file.Status()
	file.MarkCompleted(storageURL)
	url, err := file.GenerateDownloadURL() // err if file has not completed upload

	// errors and retries:
	// what happens when aborted?
	// what happens when part fails?
	// what happens when server restarts?

}
