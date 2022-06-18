package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/baguss42/machtwatch/entity"
)

type TransactionDetailRepositoryInterface interface {
	Create(context.Context, *sql.Tx, entity.TransactionDetail) error
	Get(context.Context, int64) ([]entity.TransactionDetail, error)
}

type TransactionDetailRepository struct {
	DB *sql.DB
}

func NewTransactionDetailRepository(db *sql.DB) *TransactionDetailRepository {
	return &TransactionDetailRepository{
		DB: db,
	}
}

func (t *TransactionDetailRepository) Create(ctx context.Context, tx *sql.Tx, trxDetail entity.TransactionDetail) error {
	query := fmt.Sprintf("INSERT INTO transaction_details(transaction_id, product_id, price, price_reduction, quantity, final_price) VALUES (%v, %v, %v, %v, %v, %v)", trxDetail.TransactionID, trxDetail.ProductID, trxDetail.Price, trxDetail.PriceReduction, trxDetail.Quantity, trxDetail.FinalPrice)

	_, err := tx.ExecContext(ctx, query)

	return err
}

func (t *TransactionDetailRepository) Get(ctx context.Context, transactionID int64) (result []entity.TransactionDetail, err error) {
	query := fmt.Sprintf("SELECT * FROM transaction_details WHERE transaction_id = %v", transactionID)

	rows, err := t.DB.QueryContext(ctx, query)
	defer rows.Close()

	for rows.Next() {
		td := entity.TransactionDetail{}
		err = rows.Scan(&td.ID, &td.ProductID, &td.Price, &td.PriceReduction, &td.Quantity, &td.Quantity, &td.FinalPrice, &td.CreatedAt)
		if err != nil {
			return
		}
		result = append(result, td)
	}

	return
}