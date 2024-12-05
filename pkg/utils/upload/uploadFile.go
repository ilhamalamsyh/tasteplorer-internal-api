package utils_upload_file

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func UploadToCloudinary(ctx context.Context, file *multipart.FileHeader, folder string) (*uploader.UploadResult, error) {

	environment := os.Getenv("PROJECT_ENV")
	devParentPath := os.Getenv("DEVELOPMENT_PARENT_PATH")
	prodParentPath := os.Getenv("PRODUCTION_PARENT_PATH")

	cloudinaryCloudName := []byte(os.Getenv("CLOUDINARY_CLOUD_NAME"))
	cloudinaryAPIKey := []byte(os.Getenv("CLOUDINARY_API_KEY"))
	cloudinaryAPISecret := []byte(os.Getenv("CLOUDINARY_API_SECRET"))

	cld, err := cloudinary.NewFromParams(string(cloudinaryCloudName), string(cloudinaryAPIKey), string(cloudinaryAPISecret))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Cloudinary client: %w", err)
	}

	fileContent, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file: %w", err)
	}

	defer fileContent.Close()

	// Generate a new filename using UUID
	renameFile := uuid.New().String()

	var uploadParams uploader.UploadParams

	if environment == "production" {
		uploadParams = uploader.UploadParams{
			PublicID: prodParentPath + "/" + folder + "/" + renameFile,
		}
	} else {
		// Upload to cloudinary
		uploadParams = uploader.UploadParams{
			PublicID: devParentPath + "/" + folder + "/" + renameFile,
		}
	}

	uploadResult, err := cld.Upload.Upload(ctx, fileContent, uploadParams)
	if err != nil {
		return nil, fmt.Errorf("failed to upload to Cloudinary: %w", err)
	}

	return uploadResult, nil
}
