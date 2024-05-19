package handler

import (
	"context"
	"exchange_api/internal"
	"exchange_api/internal/form"
	mailer_pb "exchange_api/proto/mailer_pb"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const statusSubscriptionExists = "exists"

type SubscribeHandler struct {
	subscribeClient mailer_pb.MailerServiceClient
}

func NewSubscribeHandler(subscribeClient mailer_pb.MailerServiceClient) *SubscribeHandler {
	return &SubscribeHandler{subscribeClient: subscribeClient}
}

func (h *SubscribeHandler) SubscribeUser(c *gin.Context) {
	var subscribe form.Subscribe

	if err := c.ShouldBindJSON(&subscribe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pair := &mailer_pb.SubscribeForm{Email: subscribe.Email, From: subscribe.From, To: subscribe.To}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	status, err := h.subscribeClient.SubscribeUser(ctx, pair)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if status.Status == statusSubscriptionExists {
		c.JSON(http.StatusConflict, gin.H{
			"error": status.Status,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status.Status,
	})
}

func (h *SubscribeHandler) SubscribeUserByUSD2UAH(c *gin.Context) {
	var subscribe form.LightSubscribe

	if err := c.ShouldBindJSON(&subscribe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pair := &mailer_pb.SubscribeForm{Email: subscribe.Email, From: internal.Currency_USD, To: internal.Currency_UAH}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	status, err := h.subscribeClient.SubscribeUser(ctx, pair)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if status.Status == statusSubscriptionExists {
		c.JSON(http.StatusConflict, gin.H{
			"error": status.Status,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "E-mail додано",
	})
}
