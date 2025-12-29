package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/99designs/gqlgen/graphql"
)

// UploadToCloudinary uploads a file to Cloudinary and returns the secure URL.
// It enforces a file size limit of 10MB.
func UploadToCloudinary(file graphql.Upload, folder string) (string, error) {
	// Check file size (10MB limit)
	const maxFileSize = 10 * 1024 * 1024
	if file.Size > maxFileSize {
		return "", fmt.Errorf("file size exceeds 10MB limit")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return "", fmt.Errorf("failed to initialize cloudinary: %v", err)
	}

	uploadResult, err := cld.Upload.Upload(ctx, file.File, uploader.UploadParams{
		Folder: folder,
		PublicID: file.Filename, // Optional: use filename as public ID
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	return uploadResult.SecureURL, nil
}
