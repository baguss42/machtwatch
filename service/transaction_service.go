package service

import (
	"context"
	"database/sql"
	"github.com/baguss42/machtwatch/entity"
	"github.com/baguss42/machtwatch/repository"
)

type TransactionServiceInterface interface {
	Create(context.Context, entity.TransactionOrder) error
	Get(ctx context.Context, int642 int64) ([]entity.TransactionDetail, error)
}

type TransactionService struct {
	TrxRepository repository.TransactionRepositoryInterface
	TrxDetailRepository repository.TransactionDetailRepositoryInterface
	ProductRepository repository.ProductRepositoryInterface
}

func NewTransactionService(db *sql.DB) *TransactionService {
	return &TransactionService{
		TrxRepository: repository.NewTransactionRepository(db),
		ProductRepository: repository.NewProductRepository(db),
		TrxDetailRepository: repository.NewTransactionDetailRepository(db),
	}
}

func (t *TransactionService) Create(ctx context.Context, transactionOrder entity.TransactionOrder) error {
	ctx, cancel := context.WithTimeout(ctx, *dbDuration)
	defer cancel()

	return t.TrxRepository.Create(ctx, transactionOrder)
}

func (t *TransactionService) Get(ctx context.Context, transactionID int64) ([]entity.TransactionDetail, error) {
	ctx, cancel := context.WithTimeout(ctx, *dbDuration)
	defer cancel()

	return t.TrxDetailRepository.Get(ctx, transactionID)
}