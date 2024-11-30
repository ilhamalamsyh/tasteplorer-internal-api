package banner_model

import (
	"time"
)

type Banner struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Image     string     `json:"image"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// func (b *Banner) MarshalJSON() ([]byte, error) {
// 	type Alias Banner
// 	if b.DeletedAt == nil {
// 		b.DeletedAt = nil // Explicitly set nil if you want to omit in the response
// 	}
// 	return json.Marshal((*Alias)(b))
// }
