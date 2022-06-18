package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/baguss42/machtwatch/entity"
)

type BrandRepositoryInterface interface {
	Create(ctx context.Context, brand entity.Brand) error
}

type BrandRepository struct {
	DB *sql.DB
}

func NewBrandRepository(db *sql.DB) *BrandRepository {
	return &BrandRepository{
		DB: db,
	}
}

func (b *BrandRepository) Create(ctx context.Context, brand entity.Brand) error {
	query := fmt.Sprintf("INSERT INTO Products(name, description, logo, level) VALUES ('%s', '%s', '%s', '%s')", brand.Name, brand.Description, brand.Logo, brand.Level)

	_, err := b.DB.ExecContext(ctx, query)

	return err
}