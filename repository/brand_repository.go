package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/baguss42/machtwatch/entity"
)

type BrandRepositoryInterface interface {
	Create(ctx context.Context, brand entity.Brand) entity.CustomError
}

type BrandRepository struct {
	DB *sql.DB
}

func NewBrandRepository(db *sql.DB) *BrandRepository {
	return &BrandRepository{
		DB: db,
	}
}

func (b *BrandRepository) Create(ctx context.Context, brand entity.Brand) (err entity.CustomError) {
	query := fmt.Sprintf("INSERT INTO brands(name, description, logo, level) VALUES ('%s', '%s', '%s', '%s')", brand.Name, brand.Description, brand.Logo, brand.Level)

	_, err.Err = b.DB.ExecContext(ctx, query)
	err.BuildSQLError("create")

	return
}