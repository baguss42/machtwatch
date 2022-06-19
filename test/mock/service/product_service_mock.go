package service

import (
	"context"
	"github.com/baguss42/machtwatch/entity"
	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (p *ProductServiceMock) Create(ctx context.Context, brand entity.Product) (err entity.CustomError) {
	args := p.Called(ctx, brand)
	return args.Get(0).(entity.CustomError)
}

func (p *ProductServiceMock) Get(ctx context.Context, id int64) (entity.Product, entity.CustomError) {
	args := p.Called(ctx, id)
	return args.Get(0).(entity.Product), args.Get(1).(entity.CustomError)
}
