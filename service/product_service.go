package service

import (
	"context"
	"database/sql"
	"github.com/baguss42/machtwatch/entity"
	"github.com/baguss42/machtwatch/repository"
)

type ProductServiceInterface interface {
	Create(context.Context, entity.Product) entity.CustomError
	Get(context.Context, int64) (entity.Product, entity.CustomError)
}

type ProductService struct {
	Repository repository.ProductRepositoryInterface
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{
		Repository: repository.NewProductRepository(db),
	}
}

func (p *ProductService) Create(ctx context.Context, product entity.Product) (err entity.CustomError) {
	err = entity.NewCustomError()
	if err.Err = product.Validate(); err.Err != nil {
		err.ErrorBadRequest(nil)
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, *dbDuration)
	defer cancel()

	select {
	case <-ctx.Done():
		return
	default:
		err = p.Repository.Create(ctx, product)
	}

	return
}

func (p *ProductService) Get(ctx context.Context, id int64) (product entity.Product, err entity.CustomError) {
	err = entity.NewCustomError()
	ctx, cancel := context.WithTimeout(ctx, *dbDuration)
	defer cancel()

	select {
	case <-ctx.Done():
		return
	default:
		product, err = p.Repository.Get(ctx, id)
	}

	return
}
