package dto

type Response struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Error      any    `json:"error"`
	Data       any    `json:"data"`
}
