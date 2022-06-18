package service

import (
	"context"
	"database/sql"
	"github.com/baguss42/machtwatch/entity"
	"github.com/baguss42/machtwatch/repository"
)

type BrandServiceInterface interface {
	Create(context.Context, entity.Brand) entity.CustomError
}

type BrandService struct {
	Repository repository.BrandRepositoryInterface
}

func NewBrandService(db *sql.DB) *BrandService {
	return &BrandService{
		Repository: repository.NewBrandRepository(db),
	}
}

func (b *BrandService) Create(ctx context.Context, brand entity.Brand) (err entity.CustomError) {
	err = entity.NewCustomError()
	if err.Err = brand.Validate(); err.Err != nil {
		err.ErrorBadRequest(nil)
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, *dbDuration)
	defer cancel()

	err = b.Repository.Create(ctx, brand)
	return
}