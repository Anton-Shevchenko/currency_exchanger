package jobs

import (
	"context"
	"exchange_mailer/internal"
	"exchange_mailer/pkg/mailer"
	mailer_pb "exchange_mailer/proto/exchanger_pb"
	"fmt"
	"log"
	"time"
)

type SendCurrencyJob struct {
	mailClient *mailer.Client
	repo       internal.RepoInterface
	exchanger  mailer_pb.ExchangerServiceClient
}

func NewSendCurrencyJob(
	mailClient *mailer.Client,
	repo internal.RepoInterface,
	exchanger mailer_pb.ExchangerServiceClient,
) *SendCurrencyJob {
	return &SendCurrencyJob{
		mailClient: mailClient,
		repo:       repo,
		exchanger:  exchanger,
	}
}

func (uc *SendCurrencyJob) Run() {
	users := uc.repo.GetAll()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, user := range users {
		rate, err := uc.exchanger.ExchangeByPair(ctx, &mailer_pb.CurrencyPair{
			From: user.From,
			To:   user.To,
		})

		if err != nil {
			log.Printf("Email Job Error: %v", err)
			return
		}

		go uc.mailClient.SendEmail(user.Email, fmt.Sprintf("Your rate: %s : %s - %f", user.From, user.To, rate.Rate))
	}
}
