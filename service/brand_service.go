package service

import (
	"context"
	"database/sql"
	"github.com/baguss42/machtwatch/entity"
	"github.com/baguss42/machtwatch/repository"
)

type BrandServiceInterface interface {
	Create(context.Context, entity.Brand) error
}

type BrandService struct {
	Repository repository.BrandRepositoryInterface
}

func NewBrandService(db *sql.DB) *BrandService {
	return &BrandService{
		Repository: repository.NewBrandRepository(db),
	}
}

func (b *BrandService) Create(ctx context.Context, brand entity.Brand) error {
	ctx, cancel := context.WithTimeout(ctx, *dbDuration)
	defer cancel()
	return b.Repository.Create(ctx, brand)
}