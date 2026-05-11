package core_http_response

type ErrorResponse struct {
	Message string `json:"message" example:"Human readable error message"`
	Error   string `json:"error" example:"Error description for developers"`
}
