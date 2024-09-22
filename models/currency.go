package models

type Input struct {
	Base    string
	Convert string
	Input   string
}
type APIval struct {
	From_Currency string  `json:"fromCurrency"`
	To_Currency   string  `json:"toCurrency"`
	Rate          float64 `json:"rate"`
}
type Output struct {
	Converstion float64
	Value       float64
}
