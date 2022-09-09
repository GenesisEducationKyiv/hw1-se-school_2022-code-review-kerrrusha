package model

type ErrorResponse struct {
	Error string
}
type SuccessResponse struct {
	Success string
}
type RateResponse struct {
	Time           string
	Asset_id_base  string
	Asset_id_quote string
	Rate           float64
}
type RateValue struct {
	Rate uint32 `json:"rate"`
}
