package handler

import (
	"context"
	proto "exchange_mailer/proto/mailer_pb"
	"fmt"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// should be in transaction
func TestCreateSubscriber(t *testing.T) {
	conn, err := grpc.NewClient("air-mailer:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("the connection with the server cannot be established")
	}
	defer conn.Close()

	client := proto.NewMailerServiceClient(conn)

	request := &proto.SubscribeForm{
		From:  "UAH",
		To:    "USD",
		Email: "test@mail.ca",
	}

	res, err := client.SubscribeUser(context.Background(), request)
	if err != nil {
		t.Fatalf("CREATE FAILED: %v", err)
	}

	fmt.Println(res.Status)
}
