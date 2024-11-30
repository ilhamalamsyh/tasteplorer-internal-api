package employee_dto

import "time"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type EmployeeDto struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"created_at"` // Timestamp when the row was created
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterDto struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (LoginRequest) CustomMessagesValidationX() map[string]string {
	return map[string]string{
		"Email.required":    "The email field is required.",
		"Email.email":       "The email field must be email format.",
		"Password.required": "The password field is required.",
		"Password.min":      "The password character must be greater than 8.",
	}
}
