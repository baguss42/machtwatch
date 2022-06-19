package service

import (
	"context"
	"github.com/baguss42/machtwatch/entity"
	"github.com/stretchr/testify/mock"
)

type TransactionServiceMock struct {
	mock.Mock
}

func (p *TransactionServiceMock) Create(ctx context.Context, brand entity.TransactionOrder) (err entity.CustomError) {
	args := p.Called(ctx, brand)
	return args.Get(0).(entity.CustomError)
}

func (p *TransactionServiceMock) Get(ctx context.Context, id int64) ([]entity.TransactionDetail, entity.CustomError) {
	args := p.Called(ctx, id)
	return args.Get(0).([]entity.TransactionDetail), args.Get(1).(entity.CustomError)
}
