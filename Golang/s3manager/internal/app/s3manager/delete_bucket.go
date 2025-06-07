package s3manager

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HandleDeleteBucket(s3 S3) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bucketName := mux.Vars(r)["bucketName"]

		err := s3.RemoveBucket(r.Context(), bucketName)
		if err != nil {
			handleHTTPError(w, fmt.Errorf("error removing bucket: %w", err))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
