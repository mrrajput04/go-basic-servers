package utils

import (
	"context"
	"errors"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// UTC uploads an image file to Cloudinary and returns the secure URL (likely a misnamed function)
func UTC(file multipart.File) (string, error) {
	urlCloudinary := os.Getenv("CLOUDINARY_URL")
	cloudService, err := cloudinary.NewFromURL(urlCloudinary)
	if err != nil {
		return "", errors.New("failed to create Cloudinary service" + err.Error())
	}

	ctx := context.Background()
	resp, err := cloudService.Upload.Upload(ctx, file, uploader.UploadParams{})
	if err != nil {
		return "", errors.New("failed to uplaod image to cloudinary" + err.Error())
	}

	return resp.SecureURL, nil
}
