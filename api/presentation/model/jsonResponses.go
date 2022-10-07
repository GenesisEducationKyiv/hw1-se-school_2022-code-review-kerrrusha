package model

type ErrorResponse struct {
	Error string `json:"error"`
}
type SuccessResponse struct {
	Success string `json:"success"`
}
type RateValue struct {
	Rate uint32 `json:"rate"`
}
