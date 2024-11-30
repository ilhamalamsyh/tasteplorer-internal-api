package response_dto

type ResponseDto struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}
