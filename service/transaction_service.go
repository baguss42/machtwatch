package service

import (
	"context"
	"database/sql"
	"github.com/baguss42/machtwatch/entity"
	"github.com/baguss42/machtwatch/repository"
)

type TransactionServiceInterface interface {
	Create(context.Context, entity.TransactionOrder) entity.CustomError
	Get(ctx context.Context, int642 int64) ([]entity.TransactionDetail, entity.CustomError)
}

type TransactionService struct {
	TrxRepository       repository.TransactionRepositoryInterface
	TrxDetailRepository repository.TransactionDetailRepositoryInterface
	ProductRepository   repository.ProductRepositoryInterface
}

func NewTransactionService(db *sql.DB) *TransactionService {
	return &TransactionService{
		TrxRepository:       repository.NewTransactionRepository(db),
		ProductRepository:   repository.NewProductRepository(db),
		TrxDetailRepository: repository.NewTransactionDetailRepository(db),
	}
}

func (t *TransactionService) Create(ctx context.Context, transactionOrder entity.TransactionOrder) (err entity.CustomError) {
	ctx, cancel := context.WithTimeout(ctx, *dbDuration)
	defer cancel()

	select {
	case <-ctx.Done():
		return
	default:
		err = t.TrxRepository.Create(ctx, transactionOrder)
	}

	return
}

func (t *TransactionService) Get(ctx context.Context, transactionID int64) (result []entity.TransactionDetail, err entity.CustomError) {
	ctx, cancel := context.WithTimeout(ctx, *dbDuration)
	defer cancel()

	select {
	case <-ctx.Done():
		return
	default:
		result, err = t.TrxDetailRepository.Get(ctx, transactionID)
	}

	if len(result) < 1 {
		err.ErrorNotFound()
	}

	return result, err
}
