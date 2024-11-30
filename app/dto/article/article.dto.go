package article_dto

import "time"

type ArticleDto struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ImageUrl    string     `json:"image_url"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type ArticleRequestDto struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=20"`
	ImageUrl    string `json:"image_url" validate:"required,url"`
}

func (ArticleRequestDto) CustomMessagesValidation() map[string]string {
	return map[string]string{
		"Title.required":       "The title field is required.",
		"Title.min":            "The title must be at least 3 characters long.",
		"Title.max":            "The title cannot exceed 100 characters.",
		"Description.required": "The description field is required.",
		"Description.min":      "The description must be at least 20 characters long.",
		"ImageUrl.required":    "The image_url is required.",
		"ImageUrl.url":         "The image_url must be a valid URL.",
	}
}
