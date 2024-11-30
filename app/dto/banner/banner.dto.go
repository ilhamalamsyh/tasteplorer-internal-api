package banner_dto

import "time"

type BannerDto struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Image     string     `json:"image"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type BannerRequestDto struct {
	Title string `json:"title" validate:"required,min=3,max=100"`
	Image string `json:"image" validate:"required,url"`
}

func (BannerRequestDto) CustomMessagesValidation() map[string]string {
	return map[string]string{
		"Title.required": "The title field is required.",
		"Title.min":      "The title must be at least 3 characters long.",
		"Title.max":      "The title cannot exceed 100 characters.",
		"Image.required": "The Image is required.",
		"Image.url":      "The Image must be a valid URL.",
	}
}
