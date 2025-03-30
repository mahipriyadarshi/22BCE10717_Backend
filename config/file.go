package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"file-sharing-app/config"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Limit upload size to 10MB

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	bucket := os.Getenv("S3_BUCKET")
	key := header.Filename

	// Upload file to S3
	_, err = config.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	})
	if err != nil {
		http.Error(w, "Failed to upload file", http.StatusInternalServerError)
		return
	}

	s3URL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "File uploaded successfully",
		"url":     s3URL,
	})
}
