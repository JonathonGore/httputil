package httputil

// SuccessResponse is a JSON wrapper for a successful operation.
type SuccessResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// ErrorResponse is a JSON wrapper for an unsuccessful operation.
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
