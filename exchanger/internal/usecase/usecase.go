package usecase

import (
	"encoding/json"
	"exchange_exchanger/config"
	"exchange_exchanger/internal/form"
	"exchange_exchanger/pkg/cache"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const rateTTLInMinute = 5

type ExchangeUseCase struct {
	store cache.Cache
	cnf   *config.Config
}

func NewExchangeUseCase(store cache.Cache, cnf *config.Config) *ExchangeUseCase {
	return &ExchangeUseCase{store: store, cnf: cnf}
}

func (c *ExchangeUseCase) GetCurrencyByPair(currencyPair *form.CurrencyPair) (*form.Rate, error) {
	cacheKey := fmt.Sprintf("%s:%s", currencyPair.From, currencyPair.To)
	rate, isExists := c.store.Get(cacheKey)

	if isExists == true {
		value, err := strconv.ParseFloat(rate, 32)

		if err != nil {
			return nil, err
		}

		return &form.Rate{Value: float32(value)}, nil
	}
	fmt.Println(fmt.Sprintf("%s/pair/%s/%s",
		c.cnf.CurrencyApiUrl,
		currencyPair.From,
		currencyPair.To,
	))
	resp, err := http.Get(fmt.Sprintf("%s/pair/%s/%s",
		c.cnf.CurrencyApiUrl,
		currencyPair.From,
		currencyPair.To,
	))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var rateEntity form.Rate
	err = json.Unmarshal(body, &rateEntity)

	if err != nil {
		return nil, err
	}

	c.store.Set(cacheKey, rateEntity.Value, time.Minute*rateTTLInMinute)

	return &rateEntity, nil
}
