package service

import (
	"context"
	"database/sql"
	"github.com/baguss42/machtwatch/entity"
	"github.com/baguss42/machtwatch/repository"
	"github.com/baguss42/machtwatch/test/factory"
	mock "github.com/baguss42/machtwatch/test/mock/repository"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransactionServiceTestSuite struct {
	suite.Suite
	mockTransactionRepository       *mock.TransactionRepositoryMock
	mockTransactionDetailRepository *mock.TransactionDetailRepositoryMock
	serviceInstance                 TransactionService
	transactionOrder                entity.TransactionOrder
	transactionDetails              []entity.TransactionDetail
	customError                     entity.CustomError
}

func (ts *TransactionServiceTestSuite) SetupTest() {
	ts.mockTransactionRepository = &mock.TransactionRepositoryMock{}
	ts.mockTransactionDetailRepository = &mock.TransactionDetailRepositoryMock{}
	ts.serviceInstance = TransactionService{TrxRepository: ts.mockTransactionRepository, TrxDetailRepository: ts.mockTransactionDetailRepository}
	ts.transactionOrder = (&factory.TransactionOrderFactory{}).Build()
	ts.transactionDetails = (&factory.TransactionDetailsFactory{}).Build()
	ts.customError = entity.NewCustomError()
	InitDBTimeOut()
}

func (ts *TransactionServiceTestSuite) TestNewTransactionService() {
	var db *sql.DB
	transactionService := NewTransactionService(db)
	ts.Equal(&TransactionService{
		TrxRepository: &repository.TransactionRepository{
			DB:                          db,
			TransactionDetailRepository: &repository.TransactionDetailRepository{DB: db},
			ProductRepository:           &repository.ProductRepository{DB: db},
		},
		ProductRepository:   &repository.ProductRepository{DB: db},
		TrxDetailRepository: &repository.TransactionDetailRepository{DB: db},
	}, transactionService)
}

func (ts *TransactionServiceTestSuite) TestCreate() {
	testCases := []struct {
		name             string
		transactionOrder entity.TransactionOrder
		err              entity.CustomError
	}{
		{
			name:             "Test case 1: success create transaction | 200",
			transactionOrder: ts.transactionOrder,
			err:              ts.customError,
		},
	}

	for _, tc := range testCases {
		ctx, cancel := context.WithTimeout(context.Background(), *dbDuration)

		ts.mockTransactionRepository.On("Create", ctx, tc.transactionOrder).Return(tc.err).Once()
		err := ts.serviceInstance.Create(ctx, tc.transactionOrder)

		ts.Nil(err.Err)

		cancel()
	}
}

func (ts *TransactionServiceTestSuite) TestGet() {
	testCases := []struct {
		name               string
		id                 int64
		transactionDetails []entity.TransactionDetail
		err                entity.CustomError
	}{
		{
			name:               "Test case 1: success get product | 200",
			transactionDetails: ts.transactionDetails,
			id:                 1,
			err:                ts.customError,
		},
	}

	for _, tc := range testCases {
		ctx, cancel := context.WithTimeout(context.Background(), *dbDuration)

		ts.mockTransactionDetailRepository.On("Get", ctx, tc.id).Return(tc.transactionDetails, tc.err).Once()
		_, err := ts.serviceInstance.Get(ctx, tc.id)

		ts.Nil(err.Err)

		cancel()
	}
}

func TestTransactionService(t *testing.T) {
	suite.Run(t, new(TransactionServiceTestSuite))
}
