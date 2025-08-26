package file

import (
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Interact with AWS S3
type S3BlobStore struct {
	s3 *s3.Client
}

func NewBlobStore(s3 *s3.Client) *S3BlobStore {
	return &S3BlobStore{s3: s3}
}

// S3 UploadId
type UploadId string

func (b *S3BlobStore) CreateMultiPartUpload(file File) (UploadId, error) {
	return "1", nil
}

// Info that browser needs to know to upload
type UploadInfo struct {
	PartNumber int
	ChunkSize  int
	UploadURL  url.URL
	Expiry     time.Time
}

func (b *S3BlobStore) GeneratePresignedUploadURLs(file File) ([]UploadInfo, error) {
	return []UploadInfo{
		UploadInfo{},
	}, nil
}

func (b *S3BlobStore) CompleteMultiPartUpload(file File) error {
	return nil
}
