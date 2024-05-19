package main

import (
	"exchange_mailer/config"
	database "exchange_mailer/internal/db"
	"exchange_mailer/internal/grpc/handler"
	"exchange_mailer/internal/jobs"
	"exchange_mailer/internal/models"
	repo "exchange_mailer/internal/repository"
	"exchange_mailer/internal/usecase"
	"exchange_mailer/pkg/mailer"
	"exchange_mailer/pkg/scheduler"
	exchanger_pb "exchange_mailer/proto/exchanger_pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
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
		log.Fatalf("Error starting: %v", err)
	}
	db := database.DbConn()
	migrations(db)

	exchangerClient, err := grpc.NewClient(cnf.ExchangerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}

	grpcServer := grpc.NewServer()
	subscriptionRepo := repo.NewSubscriptionRepository(db)

	subscribeUseCase := usecase.NewSubscribeUseCase(subscriptionRepo)

	defer exchangerClient.Close()

	exchangerService := exchanger_pb.NewExchangerServiceClient(exchangerClient)
	handler.NewMailerHandler(grpcServer, subscribeUseCase)

	mailerClient := mailer.NewClient(mailer.Config{
		PrivKey:     cnf.MailPrivKey,
		PubKey:      cnf.MailPubKey,
		SenderEmail: cnf.MailSender,
	})
	job := jobs.NewSendCurrencyJob(mailerClient, subscriptionRepo, exchangerService)
	scheduler.NewScheduler(job).Daily()

	log.Fatal(grpcServer.Serve(lis))
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.Subscriber{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}
