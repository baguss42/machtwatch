package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/baguss42/machtwatch/entity"
)

type ProductRepositoryInterface interface {
	Create(context.Context, entity.Product) error
	Get(context.Context, int64) (entity.Product, error)
	UpdateStock(context.Context, *sql.Tx, int64, int) error
}

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (p *ProductRepository) Create(ctx context.Context, product entity.Product) error {
	query := fmt.Sprintf("INSERT INTO products(brand_id, title, description, price, price_reduction, stock) " +
		"VALUES (%v, '%s', '%s', %v, %v, %v)",
		product.BrandID, product.Title, product.Description, product.Price, product.PriceReduction, product.Stock)

	_, err := p.DB.ExecContext(ctx, query)

	return err
}

func (p *ProductRepository) Get(ctx context.Context, id int64) (result entity.Product, err error) {
	query := fmt.Sprintf("SELECT * FROM products WHERE id = %v", id)

	err = p.DB.QueryRowContext(ctx, query).Scan(&result.ID, &result.BrandID, &result.Title,
		&result.Description, &result.Price, &result.PriceReduction, &result.Stock,
		&result.IsActive, &result.CreatedAt)
	if err != nil {
		return
	}

	return
}

func (p *ProductRepository) UpdateStock(ctx context.Context, tx *sql.Tx, id int64, stock int) error {
	query := fmt.Sprintf("UPDATE products SET stock = %v WHERE id = %v", stock, id)

	_, err := tx.ExecContext(ctx, query)

	return err
}