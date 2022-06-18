package service

import (
	"context"
	"database/sql"
	"github.com/baguss42/machtwatch/entity"
	"github.com/baguss42/machtwatch/repository"
)

type ProductServiceInterface interface {
	Create(context.Context, entity.Product) error
	Get(context.Context, int64) (entity.Product, error)
}

type ProductService struct {
	Repository repository.ProductRepositoryInterface
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{
		Repository: repository.NewProductRepository(db),
	}
}

func (p *ProductService) Create(ctx context.Context, Product entity.Product) error {
	ctx, cancel := context.WithTimeout(ctx, *dbDuration)
	defer cancel()
	return p.Repository.Create(ctx, Product)
}

func (p *ProductService) Get(ctx context.Context, id int64) (entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, *dbDuration)
	defer cancel()
	return p.Repository.Get(ctx, id)
}