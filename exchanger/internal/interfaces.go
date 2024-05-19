package internal

import (
	"exchange_exchanger/internal/form"
)

type ExchangeUseCaseInterface interface {
	GetCurrencyByPair(currencyPair *form.CurrencyPair) (*form.Rate, error)
}
