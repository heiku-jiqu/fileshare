package file

import (
	"time"

	"github.com/heiku-jiqu/fileshare/user"
)

// File tracks the progress of a file upload.
type File struct {
	Status           FileStatus
	Name             string // name of the file
	size             int64
	ownerId          user.UserId
	sharedWith       []user.UserId
	createdTimestamp time.Time

	// S3 related fields
	parts             []Part
	storageURL        string // where the actual data is stored at
	multipartUploadId string // Amazon S3 multipart upload ID

	// checksum         []byte
	// checksumType     checksumType
}

type FileStatus string

const Started FileStatus = "STARTED"                   // Initial state, waiting for data to be uploaded
const AllPartsUploaded FileStatus = "ALLPARTSUPLOADED" // State when all file parts have been uploaded but not merged
const Completed FileStatus = "COMPLETED"               // Final success state when all parts are merged, ready to be downloaded
const TimedOut FileStatus = "TIMEDOUT"                 // Error state when upload window is over but file not completed uploaded
const Aborted FileStatus = "ABORTED"                   // Error state when upload has been aborted at s3

type Part struct {
	number    int    // the part's order number
	uploaded  bool   // whether the part has been uploaded successfully
	etag      string // the s3 etag value of the part
	byteStart int64
	byteEnd   int64
	chunkSize int64 // chunkSize = (byteEnd - byteStart + 1)

	// checksum     []byte // the part's checksum using checksumType algorithm
	// checksumType checksumType
}

type checksumType string

func NewFile(name string, size int64, userId user.UserId) File {
	var parts []Part
	var chunkSize int64 = 1_048_576 * 16 // 16 MB

	for i, byterange := range chunkFile(size, chunkSize) {
		parts = append(parts, Part{number: i, uploaded: false, etag: "", byteStart: byterange.startByte, byteEnd: byterange.endByte})
	}

	return File{
		Status:           Started,
		Name:             name,
		size:             size,
		parts:            parts,
		ownerId:          userId,
		sharedWith:       []user.UserId{},
		createdTimestamp: time.Now(),
		storageURL:       "",
	}
}

// Returns the unique f.owner/f.name of the file
func (f File) Key() string {
	return string(f.ownerId) + "/" + f.Name
}
func (f File) PendingParts() []Part {
	out := []Part{}
	for _, p := range f.parts {
		if p.uploaded == false {
			out = append(out, p)
		}
	}
	return out
}

func (f File) MarkPartUploaded(partNum int, etag string) File {
	if f.Status == Completed || f.Status == AllPartsUploaded || f.Status == Aborted {
		return f
	}

	for _, part := range f.parts {
		if part.number == partNum {
			part.etag = etag
			part.uploaded = true
		}
	}

	if f.isAllUploaded() {
		f.Status = Completed
	}

	return f
}

func (f File) MarkCompleted(storageURL string) File {
	f.storageURL = storageURL
	f.Status = Completed
	return f
}

func (f File) isAllUploaded() bool {
	allUploaded := true
	for _, part := range f.parts {
		if !part.uploaded {
			allUploaded = false
			break
		}
	}
	return allUploaded
}

type byteRange struct {
	startByte int64
	endByte   int64
}

func chunkFile(fileSize int64, chunkSize int64) []byteRange {
	numChunks := (fileSize + (chunkSize - 1)) / chunkSize // int division rounded up
	out := make([]byteRange, numChunks)
	var start, end int64
	for i := range numChunks {
		start += chunkSize + 1
		end = min(start+chunkSize, fileSize)
		out[i] = byteRange{start, end}
	}
	return out
}
