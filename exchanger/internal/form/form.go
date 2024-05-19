package form

type CurrencyPair struct {
	From string
	To   string
}

type Rate struct {
	Value float32 `json:"conversion_rate"`
}
