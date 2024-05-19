package handler

import (
	"context"
	"errors"
	"exchange_mailer/internal"
	"exchange_mailer/internal/models"
	mailer_pb "exchange_mailer/proto/mailer_pb"
	"google.golang.org/grpc"
)

type MailerServStruct struct {
	useCase internal.SubscribeUseCaseInterface
	mailer_pb.MailerServiceServer
}

func NewMailerHandler(grpcServer *grpc.Server, usecase internal.SubscribeUseCaseInterface) {
	mailerGrpc := &MailerServStruct{useCase: usecase}
	mailer_pb.RegisterMailerServiceServer(grpcServer, mailerGrpc)
}

func (srv *MailerServStruct) SubscribeUser(ctx context.Context, req *mailer_pb.SubscribeForm) (*mailer_pb.Response, error) {
	data := srv.transformPushRPC(req)

	if data.Email == "" || data.From == "" || data.To == "" {
		return nil, errors.New("please provide all fields")
	}
	status := srv.useCase.SubscribeNewUser(data)

	return &mailer_pb.Response{Status: status.Status}, nil
}

func (srv *MailerServStruct) transformPushRPC(req *mailer_pb.SubscribeForm) *models.Subscriber {
	return &models.Subscriber{
		Email: req.GetEmail(),
		From:  req.GetFrom(),
		To:    req.GetTo(),
	}
}
