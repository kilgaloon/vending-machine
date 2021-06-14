package http

// ErrorResponse that will be returned 
type ErrorResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
}