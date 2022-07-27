package model

type RateAnswer struct {
	Time           string
	Asset_id_base  string
	Asset_id_quote string
	Rate           float64
}
type RateValue struct {
	Rate uint32 `json:"rate"`
}
