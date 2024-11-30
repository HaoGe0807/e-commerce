package ebus

type Money struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
