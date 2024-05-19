package usecase

import (
	"exchange_mailer/internal"
	"exchange_mailer/internal/grpc"
	"exchange_mailer/internal/models"
)

type SubscribeUseCase struct {
	repo internal.RepoInterface
}

func NewSubscribeUseCase(repo internal.RepoInterface) internal.SubscribeUseCaseInterface {
	return &SubscribeUseCase{repo}
}

func (uc *SubscribeUseCase) SubscribeNewUser(form *models.Subscriber) *grpc.Response {
	exists := uc.repo.Exists(form.Email)

	// can be by unique index
	if exists {
		return &grpc.Response{
			Status: "exists",
		}
	}

	_, err := uc.repo.Create(form)

	if err != nil {
		return nil
	}

	return &grpc.Response{
		Status: "new",
	}
}
