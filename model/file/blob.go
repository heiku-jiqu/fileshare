package file

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// number of seconds that presigned URLs are valid for
const validPresignDuration = 15 * 60 // 15 minutes

// Interact with AWS S3
type S3BlobStore struct {
	s3        *s3.Client
	presigner *s3.PresignClient
	bucket    string
}

func NewBlobStore(s3client *s3.Client, bucket string) *S3BlobStore {
	return &S3BlobStore{s3: s3client, presigner: s3.NewPresignClient(s3client), bucket: bucket}
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
	UploadURL  string
	Expiry     time.Time
}

func (b *S3BlobStore) GeneratePresignedUploadURLs(file File, uploadId UploadId) ([]UploadInfo, error) {
	params := &s3.UploadPartInput{
		Bucket: aws.String(b.bucket), Key: aws.String(file.Key()), UploadId: aws.String(string(uploadId)), PartNumber: aws.Int32(1),
	}
	presignedReq, err := b.presigner.PresignUploadPart(context.TODO(), params, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(validPresignDuration)
	})
	if err != nil {
		log.Printf("Error presigning: %+v", err)
		return nil, err
	}
	return []UploadInfo{
		UploadInfo{
			PartNumber: 1,
			ChunkSize:  1,
			UploadURL:  presignedReq.URL,
			Expiry:     time.Now().Add(time.Duration(validPresignDuration)), // TODO: not very accurate
		},
	}, nil
}

func (b *S3BlobStore) CompleteMultiPartUpload(file File, uploadId UploadId) error {
	params := &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(b.bucket),
		Key:      aws.String(file.Key()),
		UploadId: aws.String(string(uploadId)),
	}
	_, err := b.s3.CompleteMultipartUpload(context.TODO(), params)
	if err != nil {
		log.Printf("Error completing multipart upload: %+v", err)
		return err
	}
	return nil
}
