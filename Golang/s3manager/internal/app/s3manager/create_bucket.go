package s3manager

import (
	"encoding/json"
	"fmt"
	"github.com/minio/minio-go/v7"
	"net/http"
)

// HandleCreateBucket creates a new bucket.
func HandleCreateBucket(s3 S3) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bucket minio.BucketInfo
		err := json.NewDecoder(r.Body).Decode(&bucket)
		if err != nil {
			handleHTTPError(w, fmt.Errorf("error decoding body JSON: %w", err))
			return
		}

		err = s3.MakeBucket(r.Context(), bucket.Name, minio.MakeBucketOptions{})
		if err != nil {
			handleHTTPError(w, fmt.Errorf("error making bucket: %w", err))
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(bucket)
		if err != nil {
			handleHTTPError(w, fmt.Errorf("error encoding JSON: %w", err))
			return
		}
	}
}
