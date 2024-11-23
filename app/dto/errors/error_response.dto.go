package dto_custom_error

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
