package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/baguss42/machtwatch/entity"
)

type TransactionDetailRepositoryInterface interface {
	Create(context.Context, *sql.Tx, entity.TransactionDetail) entity.CustomError
	Get(context.Context, int64) ([]entity.TransactionDetail, entity.CustomError)
}

type TransactionDetailRepository struct {
	DB *sql.DB
}

func NewTransactionDetailRepository(db *sql.DB) *TransactionDetailRepository {
	return &TransactionDetailRepository{
		DB: db,
	}
}

func (t *TransactionDetailRepository) Create(ctx context.Context, tx *sql.Tx, trxDetail entity.TransactionDetail) (err entity.CustomError) {
	query := fmt.Sprintf("INSERT INTO transaction_details(transaction_id, product_id, price, price_reduction, quantity, final_price) VALUES (%v, %v, %v, %v, %v, %v)", trxDetail.TransactionID, trxDetail.ProductID, trxDetail.Price, trxDetail.PriceReduction, trxDetail.Quantity, trxDetail.FinalPrice)

	_, err.Err = tx.ExecContext(ctx, query)
	err.BuildSQLError("create")

	return err
}

func (t *TransactionDetailRepository) Get(ctx context.Context, transactionID int64) (result []entity.TransactionDetail, err entity.CustomError) {
	query := fmt.Sprintf("SELECT * FROM transaction_details WHERE transaction_id = %v", transactionID)

	var rows *sql.Rows
	rows, err.Err = t.DB.QueryContext(ctx, query)
	defer rows.Close()

	for rows.Next() {
		td := entity.TransactionDetail{}
		err.Err = rows.Scan(&td.ID, &td.ProductID, &td.Price, &td.PriceReduction, &td.Quantity, &td.Quantity, &td.FinalPrice, &td.CreatedAt)
		fmt.Println(err)
		if err.Err != nil {
			return
		}
		result = append(result, td)
	}
	err.BuildSQLError("get")

	return
}