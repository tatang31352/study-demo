package s3manager

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"net/url"
	"time"
)

//go:generate moq -out mocks/s3.go -pkg mocks . S3

// S3 is a client to interact with S3 storage.
type S3 interface {
	ListBuckets(ctx context.Context) ([]minio.BucketInfo, error)
	GetObject(ctx context.Context, bucketName, objectName string, opts minio.GetObjectOptions) (*minio.Object, error)
	ListObjects(ctx context.Context, bucketName string, opts minio.ListObjectsOptions) <-chan minio.ObjectInfo
	MakeBucket(ctx context.Context, bucketName string, opts minio.MakeBucketOptions) error
	RemoveBucket(ctx context.Context, bucketName string) error
	PresignedGetObject(ctx context.Context, bucketName, objectName string, expiry time.Duration, reqParams url.Values) (*url.URL, error)
	PutObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (minio.UploadInfo, error)
	RemoveObject(ctx context.Context, bucketName, objectName string, opts minio.RemoveObjectOptions) error
}

type SSEType struct {
	Type string
	Key  string
}
