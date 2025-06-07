package s3manager_test

import (
	"context"
	"github.com/matryer/is"
	"github.com/minio/minio-go/v7"
	"net/http"
	"testing"
)

func TestHandleBucketView(t *testing.T) {
	t.Parallel()
	cases := []struct {
		it                   string
		listObjectsFunc      func(context.Context, string, minio.ListObjectsOptions) <-chan minio.ObjectInfo
		bucketName           string
		path                 string
		expectedStatusCode   int
		expectedBodyContains string
	}{
		{
			it: "renders a bucket containing a file",
			listObjectsFunc: func(context.Context, string, minio.ListObjectsOptions) <-chan minio.ObjectInfo {
				objCh := make(chan minio.ObjectInfo)
				go func() {
					objCh <- minio.ObjectInfo{Key: "FILE-NAME"}
					close(objCh)
				}()
				return objCh
			},
			bucketName:           "BUCKET-NAME",
			expectedStatusCode:   http.StatusOK,
			expectedBodyContains: "FILE-NAME",
		},
		{
			it: "renders placeholder for an empty bucket",
			listObjectsFunc: func(context.Context, string, minio.ListObjectsOptions) <-chan minio.ObjectInfo {
				objCh := make(chan minio.ObjectInfo)
				close(objCh)
				return objCh
			},
			bucketName:           "BUCKET-NAME",
			expectedStatusCode:   http.StatusOK,
			expectedBodyContains: "No objects in",
		},
		{
			it: "renders a bucket containing an archive",
			listObjectsFunc: func(context.Context, string, minio.ListObjectsOptions) <-chan minio.ObjectInfo {
				objCh := make(chan minio.ObjectInfo)
				go func() {
					objCh <- minio.ObjectInfo{Key: "archive.tar.gz"}
					close(objCh)
				}()
				return objCh
			},
			bucketName:           "BUCKET-NAME",
			expectedStatusCode:   http.StatusOK,
			expectedBodyContains: "archive",
		},
		{
			it: "renders a bucket containing an image",
			listObjectsFunc: func(context.Context, string, minio.ListObjectsOptions) <-chan minio.ObjectInfo {
				objCh := make(chan minio.ObjectInfo)
				go func() {
					objCh <- minio.ObjectInfo{Key: "FILE-NAME.png"}
					close(objCh)
				}()
				return objCh
			},
			bucketName:           "BUCKET-NAME",
			expectedStatusCode:   http.StatusOK,
			expectedBodyContains: "photo",
		},
		{
			it: "renders a bucket containing a sound file",
			listObjectsFunc: func(context.Context, string, minio.ListObjectsOptions) <-chan minio.ObjectInfo {
				objCh := make(chan minio.ObjectInfo)
				go func() {
					objCh <- minio.ObjectInfo{Key: "FILE-NAME.mp3"}
					close(objCh)
				}()
				return objCh
			},
			bucketName:           "BUCKET-NAME",
			expectedStatusCode:   http.StatusOK,
			expectedBodyContains: "music_note",
		},
		{
			it: "returns error if the bucket doesn't exist",
			listObjectsFunc: func(context.Context, string, minio.ListObjectsOptions) <-chan minio.ObjectInfo {
				objCh := make(chan minio.ObjectInfo)
				go func() {
					objCh <- minio.ObjectInfo{Err: errBucketDoesNotExist}
					close(objCh)
				}()
				return objCh
			},
			bucketName:           "BUCKET-NAME",
			expectedStatusCode:   http.StatusNotFound,
			expectedBodyContains: http.StatusText(http.StatusNotFound),
		},
		{
			it: "returns error if there is an S3 error",
			listObjectsFunc: func(context.Context, string, minio.ListObjectsOptions) <-chan minio.ObjectInfo {
				objCh := make(chan minio.ObjectInfo)
				go func() {
					objCh <- minio.ObjectInfo{Err: errS3}
					close(objCh)
				}()
				return objCh
			},
			bucketName:           "BUCKET-NAME",
			expectedStatusCode:   http.StatusInternalServerError,
			expectedBodyContains: http.StatusText(http.StatusInternalServerError),
		},
		{
			it: "renders a bucket with folder",
			listObjectsFunc: func(context.Context, string, minio.ListObjectsOptions) <-chan minio.ObjectInfo {
				objCh := make(chan minio.ObjectInfo)
				go func() {
					objCh <- minio.ObjectInfo{Key: "AFolder/"}
					close(objCh)
				}()
				return objCh
			},
			bucketName:           "BUCKET-NAME",
			expectedStatusCode:   http.StatusOK,
			expectedBodyContains: "folder",
		},
		{
			it: "renders a bucket with path",
			listObjectsFunc: func(context.Context, string, minio.ListObjectsOptions) <-chan minio.ObjectInfo {
				objCh := make(chan minio.ObjectInfo)
				close(objCh)
				return objCh
			},
			bucketName:           "BUCKET-NAME",
			path:                 "abc/def",
			expectedStatusCode:   http.StatusOK,
			expectedBodyContains: "def",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)
		})

		s3 := &mocks.S3Mock{
			ListObjectsFunc: tc.listObjectsFunc,
		}
	}

}
