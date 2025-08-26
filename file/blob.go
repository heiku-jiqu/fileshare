package file

import (
	"context"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Interact with AWS S3
type S3BlobStore struct {
	s3     *s3.Client
	bucket string
}

func NewBlobStore(s3 *s3.Client) *S3BlobStore {
	return &S3BlobStore{s3: s3}
}

// S3 UploadId
type UploadId string

func (b *S3BlobStore) CreateMultiPartUpload(file File) (UploadId, error) {
	key := file.Key()
	params := &s3.CreateMultipartUploadInput{
		Bucket: &b.bucket,
		Key:    &key,
		//CheckSum
	}
	out, err := b.s3.CreateMultipartUpload(context.TODO(), params)
	if err != nil {
		return "", err
	}
	return UploadId(*out.UploadId), nil
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
