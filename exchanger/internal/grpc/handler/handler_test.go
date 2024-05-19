package handler

import (
	"context"
	"exchange_exchanger/proto"
	"fmt"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// can be mocked
func TestCreateUser(t *testing.T) {
	conn, err := grpc.NewClient("exchanger:5002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer conn.Close()

	client := proto.NewExchangerServiceClient(conn)

	form := &proto.CurrencyPair{From: "USD", To: "UAH"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rate, err := client.ExchangeByPair(ctx, form)
	if err != nil {
		t.Fatalf("EXCHANGE FAILED: %v", err)
	}

	fmt.Println(rate.Rate)
}
