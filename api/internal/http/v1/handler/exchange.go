package handler

import (
	"context"
	"exchange_api/internal"
	"exchange_api/internal/form"
	exchanger_pb "exchange_api/proto/exchanger_pb"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CurrencyHandler struct {
	currencyService exchanger_pb.ExchangerServiceClient
}

func NewCurrencyHandler(service exchanger_pb.ExchangerServiceClient) *CurrencyHandler {
	return &CurrencyHandler{
		currencyService: service,
	}
}

func (h *CurrencyHandler) GetRateByPair(c *gin.Context) {
	var formPair form.CurrencyPair

	if err := c.ShouldBindJSON(&formPair); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rate, err := h.currencyService.ExchangeByPair(ctx, &exchanger_pb.CurrencyPair{From: formPair.From, To: formPair.To})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"type": rate.Rate,
	})
}

func (h *CurrencyHandler) GetRateByUSD2UAH(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rate, err := h.currencyService.ExchangeByPair(
		ctx,
		&exchanger_pb.CurrencyPair{From: internal.Currency_USD, To: internal.Currency_UAH},
	)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"type": rate.Rate,
	})
}
