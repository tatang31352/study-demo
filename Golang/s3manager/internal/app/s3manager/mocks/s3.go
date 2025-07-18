package mocks

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"net/url"
	"s3manager/internal/app/s3manager"
	"sync"
	"time"
)

// Ensure, that S3Mock does implement s3manager.S3.
// If this is not the case, regenerate this file with moq.
var _ s3manager.S3 = &S3Mock{}

type S3Mock struct {
	// GetObjectFunc mocks the GetObject method.
	GetObjectFunc func(ctx context.Context, bucketName string, objectName string, opts minio.GetObjectOptions) (*minio.Object, error)

	// ListBucketsFunc mocks the ListBuckets method.
	ListBucketsFunc func(ctx context.Context) ([]minio.BucketInfo, error)

	// ListObjectsFunc mocks the ListObjects method.
	ListObjectsFunc func(ctx context.Context, bucketName string, opts minio.ListObjectsOptions) <-chan minio.ObjectInfo

	// MakeBucketFunc mocks the MakeBucket method.
	MakeBucketFunc func(ctx context.Context, bucketName string, opts minio.MakeBucketOptions) error

	// PresignedGetObjectFunc mocks the PresignedGetObject method.
	PresignedGetObjectFunc func(ctx context.Context, bucketName string, objectName string, expiry time.Duration, reqParams url.Values) (*url.URL, error)

	// PutObjectFunc mocks the PutObject method.
	PutObjectFunc func(ctx context.Context, bucketName string, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (minio.UploadInfo, error)

	// RemoveBucketFunc mocks the RemoveBucket method.
	RemoveBucketFunc func(ctx context.Context, bucketName string) error

	// RemoveObjectFunc mocks the RemoveObject method.
	RemoveObjectFunc func(ctx context.Context, bucketName string, objectName string, opts minio.RemoveObjectOptions) error

	// calls tracks calls to the methods.
	calls struct {
		// GetObject holds details about calls to the GetObject method.
		GetObject []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// BucketName is the bucketName argument value.
			BucketName string
			// ObjectName is the objectName argument value.
			ObjectName string
			// Opts is the opts argument value.
			Opts minio.GetObjectOptions
		}
		// ListBuckets holds details about calls to the ListBuckets method.
		ListBuckets []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// ListObjects holds details about calls to the ListObjects method.
		ListObjects []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// BucketName is the bucketName argument value.
			BucketName string
			// Opts is the opts argument value.
			Opts minio.ListObjectsOptions
		}
		// MakeBucket holds details about calls to the MakeBucket method.
		MakeBucket []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// BucketName is the bucketName argument value.
			BucketName string
			// Opts is the opts argument value.
			Opts minio.MakeBucketOptions
		}
		// PresignedGetObject holds details about calls to the PresignedGetObject method.
		PresignedGetObject []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// BucketName is the bucketName argument value.
			BucketName string
			// ObjectName is the objectName argument value.
			ObjectName string
			// Expiry is the expiry argument value.
			Expiry time.Duration
			// ReqParams is the reqParams argument value.
			ReqParams url.Values
		}
		// PutObject holds details about calls to the PutObject method.
		PutObject []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// BucketName is the bucketName argument value.
			BucketName string
			// ObjectName is the objectName argument value.
			ObjectName string
			// Reader is the reader argument value.
			Reader io.Reader
			// ObjectSize is the objectSize argument value.
			ObjectSize int64
			// Opts is the opts argument value.
			Opts minio.PutObjectOptions
		}
		// RemoveBucket holds details about calls to the RemoveBucket method.
		RemoveBucket []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// BucketName is the bucketName argument value.
			BucketName string
		}
		// RemoveObject holds details about calls to the RemoveObject method.
		RemoveObject []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// BucketName is the bucketName argument value.
			BucketName string
			// ObjectName is the objectName argument value.
			ObjectName string
			// Opts is the opts argument value.
			Opts minio.RemoveObjectOptions
		}
	}
	lockGetObject          sync.RWMutex
	lockListBuckets        sync.RWMutex
	lockListObjects        sync.RWMutex
	lockMakeBucket         sync.RWMutex
	lockPresignedGetObject sync.RWMutex
	lockPutObject          sync.RWMutex
	lockRemoveBucket       sync.RWMutex
	lockRemoveObject       sync.RWMutex
}

// GetObject calls GetObjectFunc.
func (mock *S3Mock) GetObject(ctx context.Context, bucketName string, objectName string, opts minio.GetObjectOptions) (*minio.Object, error) {
	if mock.GetObjectFunc == nil {
		panic("S3Mock.GetObjectFunc: method is nil but S3.GetObject was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		BucketName string
		ObjectName string
		Opts       minio.GetObjectOptions
	}{
		Ctx:        ctx,
		BucketName: bucketName,
		ObjectName: objectName,
		Opts:       opts,
	}
	mock.lockGetObject.Lock()
	mock.calls.GetObject = append(mock.calls.GetObject, callInfo)
	mock.lockGetObject.Unlock()
	return mock.GetObjectFunc(ctx, bucketName, objectName, opts)
}

// GetObjectCalls gets all the calls that were made to GetObject.
// Check the length with:
//
//	len(mockedS3.GetObjectCalls())
func (mock *S3Mock) GetObjectCalls() []struct {
	Ctx        context.Context
	BucketName string
	ObjectName string
	Opts       minio.GetObjectOptions
} {
	var calls []struct {
		Ctx        context.Context
		BucketName string
		ObjectName string
		Opts       minio.GetObjectOptions
	}
	mock.lockGetObject.RLock()
	calls = mock.calls.GetObject
	mock.lockGetObject.RUnlock()
	return calls
}

// ListBuckets calls ListBucketsFunc.
func (mock *S3Mock) ListBuckets(ctx context.Context) ([]minio.BucketInfo, error) {
	if mock.ListBucketsFunc == nil {
		panic("S3Mock.ListBucketsFunc: method is nil but S3.ListBuckets was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockListBuckets.Lock()
	mock.calls.ListBuckets = append(mock.calls.ListBuckets, callInfo)
	mock.lockListBuckets.Unlock()
	return mock.ListBucketsFunc(ctx)
}

// ListBucketsCalls gets all the calls that were made to ListBuckets.
// Check the length with:
//
//	len(mockedS3.ListBucketsCalls())
func (mock *S3Mock) ListBucketsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockListBuckets.RLock()
	calls = mock.calls.ListBuckets
	mock.lockListBuckets.RUnlock()
	return calls
}

// ListObjects calls ListObjectsFunc.
func (mock *S3Mock) ListObjects(ctx context.Context, bucketName string, opts minio.ListObjectsOptions) <-chan minio.ObjectInfo {
	if mock.ListObjectsFunc == nil {
		panic("S3Mock.ListObjectsFunc: method is nil but S3.ListObjects was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		BucketName string
		Opts       minio.ListObjectsOptions
	}{
		Ctx:        ctx,
		BucketName: bucketName,
		Opts:       opts,
	}
	mock.lockListObjects.Lock()
	mock.calls.ListObjects = append(mock.calls.ListObjects, callInfo)
	mock.lockListObjects.Unlock()
	return mock.ListObjectsFunc(ctx, bucketName, opts)
}

// ListObjectsCalls gets all the calls that were made to ListObjects.
// Check the length with:
//
//	len(mockedS3.ListObjectsCalls())
func (mock *S3Mock) ListObjectsCalls() []struct {
	Ctx        context.Context
	BucketName string
	Opts       minio.ListObjectsOptions
} {
	var calls []struct {
		Ctx        context.Context
		BucketName string
		Opts       minio.ListObjectsOptions
	}
	mock.lockListObjects.RLock()
	calls = mock.calls.ListObjects
	mock.lockListObjects.RUnlock()
	return calls
}

// MakeBucket calls MakeBucketFunc.
func (mock *S3Mock) MakeBucket(ctx context.Context, bucketName string, opts minio.MakeBucketOptions) error {
	if mock.MakeBucketFunc == nil {
		panic("S3Mock.MakeBucketFunc: method is nil but S3.MakeBucket was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		BucketName string
		Opts       minio.MakeBucketOptions
	}{
		Ctx:        ctx,
		BucketName: bucketName,
		Opts:       opts,
	}
	mock.lockMakeBucket.Lock()
	mock.calls.MakeBucket = append(mock.calls.MakeBucket, callInfo)
	mock.lockMakeBucket.Unlock()
	return mock.MakeBucketFunc(ctx, bucketName, opts)
}

// MakeBucketCalls gets all the calls that were made to MakeBucket.
// Check the length with:
//
//	len(mockedS3.MakeBucketCalls())
func (mock *S3Mock) MakeBucketCalls() []struct {
	Ctx        context.Context
	BucketName string
	Opts       minio.MakeBucketOptions
} {
	var calls []struct {
		Ctx        context.Context
		BucketName string
		Opts       minio.MakeBucketOptions
	}
	mock.lockMakeBucket.RLock()
	calls = mock.calls.MakeBucket
	mock.lockMakeBucket.RUnlock()
	return calls
}

// PresignedGetObject calls PresignedGetObjectFunc.
func (mock *S3Mock) PresignedGetObject(ctx context.Context, bucketName string, objectName string, expiry time.Duration, reqParams url.Values) (*url.URL, error) {
	if mock.PresignedGetObjectFunc == nil {
		panic("S3Mock.PresignedGetObjectFunc: method is nil but S3.PresignedGetObject was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		BucketName string
		ObjectName string
		Expiry     time.Duration
		ReqParams  url.Values
	}{
		Ctx:        ctx,
		BucketName: bucketName,
		ObjectName: objectName,
		Expiry:     expiry,
		ReqParams:  reqParams,
	}
	mock.lockPresignedGetObject.Lock()
	mock.calls.PresignedGetObject = append(mock.calls.PresignedGetObject, callInfo)
	mock.lockPresignedGetObject.Unlock()
	return mock.PresignedGetObjectFunc(ctx, bucketName, objectName, expiry, reqParams)
}

// PresignedGetObjectCalls gets all the calls that were made to PresignedGetObject.
// Check the length with:
//
//	len(mockedS3.PresignedGetObjectCalls())
func (mock *S3Mock) PresignedGetObjectCalls() []struct {
	Ctx        context.Context
	BucketName string
	ObjectName string
	Expiry     time.Duration
	ReqParams  url.Values
} {
	var calls []struct {
		Ctx        context.Context
		BucketName string
		ObjectName string
		Expiry     time.Duration
		ReqParams  url.Values
	}
	mock.lockPresignedGetObject.RLock()
	calls = mock.calls.PresignedGetObject
	mock.lockPresignedGetObject.RUnlock()
	return calls
}

// PutObject calls PutObjectFunc.
func (mock *S3Mock) PutObject(ctx context.Context, bucketName string, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (minio.UploadInfo, error) {
	if mock.PutObjectFunc == nil {
		panic("S3Mock.PutObjectFunc: method is nil but S3.PutObject was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		BucketName string
		ObjectName string
		Reader     io.Reader
		ObjectSize int64
		Opts       minio.PutObjectOptions
	}{
		Ctx:        ctx,
		BucketName: bucketName,
		ObjectName: objectName,
		Reader:     reader,
		ObjectSize: objectSize,
		Opts:       opts,
	}
	mock.lockPutObject.Lock()
	mock.calls.PutObject = append(mock.calls.PutObject, callInfo)
	mock.lockPutObject.Unlock()
	return mock.PutObjectFunc(ctx, bucketName, objectName, reader, objectSize, opts)
}

// PutObjectCalls gets all the calls that were made to PutObject.
// Check the length with:
//
//	len(mockedS3.PutObjectCalls())
func (mock *S3Mock) PutObjectCalls() []struct {
	Ctx        context.Context
	BucketName string
	ObjectName string
	Reader     io.Reader
	ObjectSize int64
	Opts       minio.PutObjectOptions
} {
	var calls []struct {
		Ctx        context.Context
		BucketName string
		ObjectName string
		Reader     io.Reader
		ObjectSize int64
		Opts       minio.PutObjectOptions
	}
	mock.lockPutObject.RLock()
	calls = mock.calls.PutObject
	mock.lockPutObject.RUnlock()
	return calls
}

// RemoveBucket calls RemoveBucketFunc.
func (mock *S3Mock) RemoveBucket(ctx context.Context, bucketName string) error {
	if mock.RemoveBucketFunc == nil {
		panic("S3Mock.RemoveBucketFunc: method is nil but S3.RemoveBucket was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		BucketName string
	}{
		Ctx:        ctx,
		BucketName: bucketName,
	}
	mock.lockRemoveBucket.Lock()
	mock.calls.RemoveBucket = append(mock.calls.RemoveBucket, callInfo)
	mock.lockRemoveBucket.Unlock()
	return mock.RemoveBucketFunc(ctx, bucketName)
}

// RemoveBucketCalls gets all the calls that were made to RemoveBucket.
// Check the length with:
//
//	len(mockedS3.RemoveBucketCalls())
func (mock *S3Mock) RemoveBucketCalls() []struct {
	Ctx        context.Context
	BucketName string
} {
	var calls []struct {
		Ctx        context.Context
		BucketName string
	}
	mock.lockRemoveBucket.RLock()
	calls = mock.calls.RemoveBucket
	mock.lockRemoveBucket.RUnlock()
	return calls
}

// RemoveObject calls RemoveObjectFunc.
func (mock *S3Mock) RemoveObject(ctx context.Context, bucketName string, objectName string, opts minio.RemoveObjectOptions) error {
	if mock.RemoveObjectFunc == nil {
		panic("S3Mock.RemoveObjectFunc: method is nil but S3.RemoveObject was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		BucketName string
		ObjectName string
		Opts       minio.RemoveObjectOptions
	}{
		Ctx:        ctx,
		BucketName: bucketName,
		ObjectName: objectName,
		Opts:       opts,
	}
	mock.lockRemoveObject.Lock()
	mock.calls.RemoveObject = append(mock.calls.RemoveObject, callInfo)
	mock.lockRemoveObject.Unlock()
	return mock.RemoveObjectFunc(ctx, bucketName, objectName, opts)
}

// RemoveObjectCalls gets all the calls that were made to RemoveObject.
// Check the length with:
//
//	len(mockedS3.RemoveObjectCalls())
func (mock *S3Mock) RemoveObjectCalls() []struct {
	Ctx        context.Context
	BucketName string
	ObjectName string
	Opts       minio.RemoveObjectOptions
} {
	var calls []struct {
		Ctx        context.Context
		BucketName string
		ObjectName string
		Opts       minio.RemoveObjectOptions
	}
	mock.lockRemoveObject.RLock()
	calls = mock.calls.RemoveObject
	mock.lockRemoveObject.RUnlock()
	return calls
}
