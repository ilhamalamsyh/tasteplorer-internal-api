package upload_dto

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
)

type PathEnum string

const MaxFileSize = 1 * 1024 * 1024

var (
	ArticlePath PathEnum = PathEnum(os.Getenv("ARTICLE_PATH"))
	BannerPath  PathEnum = PathEnum(os.Getenv("BANNER_PATH"))
)

// IsValid checks if the folder value is valid
func IsValid(f string) bool {
	articlePath := os.Getenv("ARTICLE_PATH")
	bannerPath := os.Getenv("BANNER_PATH")

	validPath := []string{articlePath, bannerPath}

	for _, path := range validPath {
		if f == path {
			return true
		}
	}
	return false
}

type UploadFileDto struct {
	File *multipart.FileHeader `json:"file" form:"file"`
	Path string                `json:"path" form:"path"`
}

// Validate validates the UploadFileDTO
func (dto *UploadFileDto) Validate() error {
	// Check if the file is provided
	if dto.File == nil {
		return errors.New("file is required")
	}

	if dto.File.Size > MaxFileSize {
		return fmt.Errorf("file size exceeds the maximum of %d MB", MaxFileSize/(1024*1024))
	}

	// Check if the folder is valid
	if !IsValid(dto.Path) {
		return errors.New("path must be one of the allowed values")
	}

	return nil
}
