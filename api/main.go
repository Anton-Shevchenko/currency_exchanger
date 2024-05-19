package main

import (
	"exchange_api/config"
	"exchange_api/internal/http/v1/api_router"
	"exchange_api/internal/http/v1/handler"
	exchanger_pb "exchange_api/proto/exchanger_pb"
	mailer_pb "exchange_api/proto/mailer_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// should be in /cmd/app
func main() {
	cnf, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error: Load Config: %v", err)
	}

	mailerClient, err := grpc.NewClient(cnf.MailerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error: Mailer Client: %v", err)
	}

	exchangerClient, err := grpc.NewClient(cnf.ExchangerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error: Exchanger Client: %v", err)
	}

	defer mailerClient.Close()
	defer exchangerClient.Close()

	mailerService := mailer_pb.NewMailerServiceClient(mailerClient)
	exchangerService := exchanger_pb.NewExchangerServiceClient(exchangerClient)

	subscribeHandler := handler.NewSubscribeHandler(mailerService)
	currencyHandler := handler.NewCurrencyHandler(exchangerService)

	v1 := api_router.NewRouter(currencyHandler, subscribeHandler)
	v1.Serve()
}
