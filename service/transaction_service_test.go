package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/baguss42/machtwatch/entity"
	"github.com/baguss42/machtwatch/repository"
	"github.com/baguss42/machtwatch/test/factory"
	mock "github.com/baguss42/machtwatch/test/mock/repository"
	"github.com/stretchr/testify/suite"
	"net/http"
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
		{
			name:             "Test case 2: request is invalid | 400",
			transactionOrder: entity.TransactionOrder{},
			err: entity.CustomError{
				HttpCode: http.StatusBadRequest,
				Err:      errors.New(fmt.Sprintf(entity.ErrFieldInvalid, "brand_id")),
			},
		},
	}

	for _, tc := range testCases {
		ctx, cancel := context.WithTimeout(context.Background(), *dbDuration)

		ts.mockTransactionRepository.On("Create", ctx, tc.transactionOrder).Return(tc.err).Once()
		err := ts.serviceInstance.Create(context.Background(), tc.transactionOrder)

		ts.Equal(tc.err, err)

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
		{
			name:               "Test case 2: record is not exist | 404",
			transactionDetails: []entity.TransactionDetail{},
			err: entity.CustomError{
				HttpCode: http.StatusNotFound,
				Err:      errors.New(entity.ErrorRecordNotExist),
			},
		},
	}

	for _, tc := range testCases {
		ctx, cancel := context.WithTimeout(context.Background(), *dbDuration)

		ts.mockTransactionDetailRepository.On("Get", ctx, tc.id).Return(tc.transactionDetails, tc.err).Once()
		transactionDetails, err := ts.serviceInstance.Get(context.Background(), tc.id)

		ts.Equal(tc.transactionDetails, transactionDetails)
		ts.Equal(tc.err, err)

		cancel()
	}
}

func TestTransactionService(t *testing.T) {
	suite.Run(t, new(TransactionServiceTestSuite))
}
