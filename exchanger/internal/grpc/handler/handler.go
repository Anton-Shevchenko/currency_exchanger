package handler

import (
	"context"
	"exchange_exchanger/internal"
	"exchange_exchanger/internal/form"
	"exchange_exchanger/proto"

	"errors"
	"google.golang.org/grpc"
)

type ExchangeHandler struct {
	useCase internal.ExchangeUseCaseInterface
	proto.ExchangerServiceServer
}

func NewExchangeHandler(grpcServer *grpc.Server, usecase internal.ExchangeUseCaseInterface) {
	exchangeGrpc := &ExchangeHandler{useCase: usecase}
	proto.RegisterExchangerServiceServer(grpcServer, exchangeGrpc)
}

func (srv *ExchangeHandler) ExchangeByPair(ctx context.Context, req *proto.CurrencyPair) (*proto.ExchangeRate, error) {
	pair := srv.transformPairRPC(req)

	if pair.From == "" || pair.To == "" {
		return nil, errors.New("please provide all fields")
	}
	rate, err := srv.useCase.GetCurrencyByPair(pair)

	if err != nil {
		return nil, err
	}

	return &proto.ExchangeRate{Rate: rate.Value}, nil
}

func (srv *ExchangeHandler) transformPairRPC(req *proto.CurrencyPair) *form.CurrencyPair {
	return &form.CurrencyPair{From: req.GetFrom(), To: req.GetTo()}
}
