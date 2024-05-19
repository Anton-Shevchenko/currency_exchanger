package internal

import (
	"exchange_mailer/internal/grpc"
	"exchange_mailer/internal/models"
)

type RepoInterface interface {
	Create(form *models.Subscriber) (*models.Subscriber, error)
	GetAll() []models.Subscriber
	Exists(email string) bool
}

type SubscribeUseCaseInterface interface {
	SubscribeNewUser(form *models.Subscriber) *grpc.Response
}
