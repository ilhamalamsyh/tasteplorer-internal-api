package employee_model

import "time"

type Employee struct {
	ID        uint       `json:"id"`
	Fullname  string     `json:"fullname"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
