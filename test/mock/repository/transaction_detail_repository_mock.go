package repository

import (
	"context"
	"database/sql"
	"github.com/baguss42/machtwatch/entity"
	"github.com/stretchr/testify/mock"
)

type TransactionDetailRepositoryMock struct {
	mock.Mock
}

func (b *TransactionDetailRepositoryMock) Create(ctx context.Context, tx *sql.Tx, trxDetail entity.TransactionDetail) entity.CustomError {
	args := b.Called(ctx, tx, trxDetail)
	return args.Get(0).(entity.CustomError)
}

func (b *TransactionDetailRepositoryMock) Get(ctx context.Context, transactionID int64) (result []entity.TransactionDetail, err entity.CustomError) {
	args := b.Called(ctx, transactionID)
	return args.Get(0).([]entity.TransactionDetail), args.Get(1).(entity.CustomError)
}
