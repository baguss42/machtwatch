package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/baguss42/machtwatch/entity"
)

type ProductRepositoryInterface interface {
	Create(context.Context, entity.Product) entity.CustomError
	Get(context.Context, int64) (entity.Product, entity.CustomError)
	UpdateStock(context.Context, *sql.Tx, int64, int) entity.CustomError
}

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (p *ProductRepository) Create(ctx context.Context, product entity.Product) (err entity.CustomError) {
	query := fmt.Sprintf("INSERT INTO products(brand_id, title, description, price, price_reduction, stock) "+
		"VALUES (%v, '%s', '%s', %v, %v, %v)",
		product.BrandID, product.Title, product.Description, product.Price, product.PriceReduction, product.Stock)

	_, err.Err = p.DB.ExecContext(ctx, query)
	err.BuildSQLError("create")

	return err
}

func (p *ProductRepository) Get(ctx context.Context, id int64) (result entity.Product, err entity.CustomError) {
	query := fmt.Sprintf("SELECT * FROM products WHERE id = %v", id)

	err.Err = p.DB.QueryRowContext(ctx, query).Scan(&result.ID, &result.BrandID, &result.Title,
		&result.Description, &result.Price, &result.PriceReduction, &result.Stock,
		&result.IsActive, &result.CreatedAt)
	err.BuildSQLError("get")

	return
}

func (p *ProductRepository) UpdateStock(ctx context.Context, tx *sql.Tx, id int64, stock int) (err entity.CustomError) {
	query := fmt.Sprintf("UPDATE products SET stock = %v WHERE id = %v", stock, id)

	_, err.Err = tx.ExecContext(ctx, query)
	err.BuildSQLError("update")

	return err
}
