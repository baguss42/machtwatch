package repository

import (
	"context"
	"database/sql"
	"github.com/baguss42/machtwatch/entity"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (b *ProductRepositoryMock) Create(ctx context.Context, brand entity.Product) (err entity.CustomError) {
	args := b.Called(ctx, brand)
	return args.Get(0).(entity.CustomError)
}

func (b *ProductRepositoryMock) Get(ctx context.Context, id int64) (entity.Product, entity.CustomError) {
	args := b.Called(ctx, id)
	return args.Get(0).(entity.Product), args.Get(1).(entity.CustomError)
}

func (b *ProductRepositoryMock) UpdateStock(ctx context.Context, tx *sql.Tx, id int64, stock int) entity.CustomError {
	args := b.Called(ctx, tx, id, stock)
	return args.Get(0).(entity.CustomError)
}
