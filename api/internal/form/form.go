package form

type CurrencyPair struct {
	From string `json:"from" binding:"required"`
	To   string `json:"to" binding:"required"`
}

type Subscribe struct {
	Email string `json:"email" binding:"required"`
	CurrencyPair
}

type LightSubscribe struct {
	Email string `json:"email" binding:"required"`
}
