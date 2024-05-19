package main

import (
	"exchange_exchanger/config"
	"exchange_exchanger/internal/grpc/handler"
	"exchange_exchanger/internal/usecase"
	"exchange_exchanger/pkg/cache/redis_cache"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

// should be in /cmd/app
func main() {
	cnf, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error: Load Config: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cnf.ServicePort))

	if err != nil {
		log.Fatalf("Error: Starting exchanger: %v", err)
	}

	cacheClient := redis_cache.NewRedisCache()
	grpcServer := grpc.NewServer()

	exchangeUseCase := usecase.NewExchangeUseCase(cacheClient, cnf)
	handler.NewExchangeHandler(grpcServer, exchangeUseCase)

	log.Fatal(grpcServer.Serve(lis))
}
