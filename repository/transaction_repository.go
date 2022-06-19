package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/baguss42/machtwatch/entity"
	"sync"
)

type TransactionRepositoryInterface interface {
	Create(context.Context, entity.TransactionOrder) entity.CustomError
}

type TransactionRepository struct {
	DB                          *sql.DB
	ProductRepository           ProductRepositoryInterface
	TransactionDetailRepository TransactionDetailRepositoryInterface
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB:                          db,
		ProductRepository:           NewProductRepository(db),
		TransactionDetailRepository: NewTransactionDetailRepository(db),
	}
}

func (t *TransactionRepository) Create(ctx context.Context, transactionOrder entity.TransactionOrder) (err entity.CustomError) {
	err = entity.NewCustomError()
	var wg sync.WaitGroup
	var mu sync.Mutex
	var result sql.Result
	var tx *sql.Tx
	var trxID int64

	cartLength := len(transactionOrder.Carts)
	errs := make(chan entity.CustomError, cartLength*4)
	wg.Add(cartLength)

	tx, err.Err = t.DB.Begin()
	if err.Err != nil {
		return
	}

	query := "INSERT INTO transactions(state) VALUES ('pending')"
	result, err.Err = tx.ExecContext(ctx, query)
	if err.Err != nil {
		return
	}

	trxID, err.Err = result.LastInsertId()
	if err.Err != nil {
		return
	}

	for _, v := range transactionOrder.Carts {
		go func(c entity.Carts, tcx *sql.Tx) {
			mu.Lock()
			var errC entity.CustomError
			defer wg.Done()
			defer mu.Unlock()

			// get product details
			product, errC := t.ProductRepository.Get(ctx, c.ProductID)
			if errC.Err != nil {
				errs <- errC
				return
			}

			if product.Stock < c.Quantity {
				errC.ErrorUnprocessableEntity(errors.New("product stock is less than quantity"))
				errs <- errC
				return
			}

			// save to transaction details
			transactionDetails := entity.TransactionDetail{
				TransactionID:  trxID,
				ProductID:      product.ID,
				Price:          product.Price,
				PriceReduction: product.PriceReduction,
				Quantity:       c.Quantity,
				FinalPrice:     (product.Price - product.PriceReduction) * int64(c.Quantity),
			}
			errC = t.TransactionDetailRepository.Create(ctx, tcx, transactionDetails)
			if errC.Err != nil {
				errs <- errC
				return
			}

			// update product reduce stock
			currentStock := product.Stock - c.Quantity
			errC = t.ProductRepository.UpdateStock(ctx, tcx, product.ID, currentStock)
			if errC.Err != nil {
				errs <- errC
				return
			}

		}(v, tx)
	}

	wg.Wait()
	close(errs)

	err = <-errs
	if err.Err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			err.Err = errRollback
			return
		}
		return
	}

	if errCommit := tx.Commit(); errCommit != nil {
		err.Err = errCommit
		return
	}

	err.BuildSQLError("create")

	return
}
