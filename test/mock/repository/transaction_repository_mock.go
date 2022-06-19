package repository

import (
	"context"
	"github.com/baguss42/machtwatch/entity"
	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (b *TransactionRepositoryMock) Create(ctx context.Context, transactionOrder entity.TransactionOrder) entity.CustomError {
	args := b.Called(ctx, transactionOrder)
	return args.Get(0).(entity.CustomError)
}
