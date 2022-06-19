package service

import (
	"context"
	"github.com/baguss42/machtwatch/entity"
	"github.com/stretchr/testify/mock"
)

type BrandServiceMock struct {
	mock.Mock
}

func (b *BrandServiceMock) Create(ctx context.Context, brand entity.Brand) (err entity.CustomError) {
	args := b.Called(ctx, brand)
	return args.Get(0).(entity.CustomError)
}
