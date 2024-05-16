package service

import (
	"bytes"
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

// SaveFile uploads the file to Cloudinary and returns the URL

func SaveFile(fileBytes []byte) (string, error) {
	// Create an io.Reader from the byte slice
	filename := uuid.New().String()
	fileReader := bytes.NewReader(fileBytes)

	cld, _ := cloudinary.NewFromParams("dsvo76qzw", "963212776811178", "iV6eu41-Z993ncaBp76UjDWposs")
	var ctx = context.Background()
	resp, err := cld.Upload.Upload(ctx, fileReader, uploader.UploadParams{ // UploadParams is a struct that contains the parameters for the upload
		PublicID: filename, // PublicID is the name of the file
	})
	if err != nil {
		log.Fatalf("Failed to upload: %v", err)
		return "", err
	}

	return resp.SecureURL, nil
}
