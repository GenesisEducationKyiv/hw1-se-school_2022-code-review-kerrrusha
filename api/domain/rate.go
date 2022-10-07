package domain

type Rate struct {
	value int
}

func CreateRate(valueFloat float64) *Rate {
	return &Rate{value: int(valueFloat)}
}

func (r *Rate) GetValue() int {
	return r.value
}
