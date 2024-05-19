package api_router

import (
	"exchange_api/internal/http/v1/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Router struct {
	currency  *handler.CurrencyHandler
	subscribe *handler.SubscribeHandler
}

func NewRouter(currency *handler.CurrencyHandler, subscribe *handler.SubscribeHandler) *Router {
	return &Router{currency: currency, subscribe: subscribe}
}

func (r *Router) Serve() {
	router := gin.Default()

	api := router.Group("/api")

	api.POST("/exchange", r.currency.GetRateByPair)
	api.GET("/rate", r.currency.GetRateByUSD2UAH)
	api.POST("/subscribe-ane", r.subscribe.SubscribeUser)
	api.POST("/subscribe", r.subscribe.SubscribeUserByUSD2UAH)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
