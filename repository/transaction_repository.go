package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/baguss42/machtwatch/entity"
	"sync"
)

type TransactionRepositoryInterface interface {
	Create(context.Context, entity.TransactionOrder) error
}

type TransactionRepository struct {
	DB *sql.DB
	ProductRepository ProductRepositoryInterface
	TransactionDetailRepository TransactionDetailRepositoryInterface
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
		ProductRepository: NewProductRepository(db),
		TransactionDetailRepository: NewTransactionDetailRepository(db),
	}
}

func (t *TransactionRepository) Create(ctx context.Context, transactionOrder entity.TransactionOrder) error {
	var wg sync.WaitGroup
	var mu sync.Mutex

	cartLength := len(transactionOrder.Carts)
	wg.Add(cartLength)
	errs := make(chan error, cartLength*4)

	tx, err := t.DB.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO transactions(state) VALUES ('pending')"
	result, err := tx.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	trxID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	for _, v := range transactionOrder.Carts {
		go func(c entity.Carts, tcx *sql.Tx) {
			mu.Lock()
			var errC error
			defer wg.Done()
			defer mu.Unlock()

			// get product details
			product, errC := t.ProductRepository.Get(ctx, c.ProductID)
			if errC != nil {
				errs <- errC
				return
			}

			if product.Stock < c.Quantity {
				errC = errors.New("quantity is greater than stock")
				errs <- errC
			}

			// save to transaction details
			transactionDetails := entity.TransactionDetail{
				TransactionID: trxID,
				ProductID: product.ID,
				Price: product.Price,
				PriceReduction: product.PriceReduction,
				Quantity: c.Quantity,
				FinalPrice: (product.Price - product.PriceReduction) * int64(c.Quantity),
			}
			errC = t.TransactionDetailRepository.Create(ctx, tcx, transactionDetails)
			if errC != nil {
				errs <- errC
				return
			}

			// update product reduce stock
			currentStock := product.Stock - c.Quantity
			errC = t.ProductRepository.UpdateStock(ctx, tcx, product.ID, currentStock)
			errs <- errC

		}(v, tx)
	}

	wg.Wait()
	close(errs)

	err = <-errs
	if err == nil {
		if errCommit := tx.Commit(); errCommit != nil {
			return errCommit
		}
	} else {
		if errRollback := tx.Rollback(); errRollback != nil {
			return errRollback
		}
		return err
	}

	return nil
}