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

type ProductServiceTestSuite struct {
	suite.Suite
	mockRepository  *mock.ProductRepositoryMock
	serviceInstance ProductService
	product         entity.Product
	customError     entity.CustomError
}

func (ts *ProductServiceTestSuite) SetupTest() {
	ts.mockRepository = &mock.ProductRepositoryMock{}
	ts.serviceInstance = ProductService{Repository: ts.mockRepository}
	ts.product = (&factory.ProductFactory{}).Build()
	ts.customError = entity.NewCustomError()
	InitDBTimeOut()
}

func (ts *ProductServiceTestSuite) TestNewProductService() {
	var db *sql.DB
	productService := NewProductService(db)
	ts.Equal(&ProductService{Repository: &repository.ProductRepository{DB: db}}, productService)
}

func (ts *ProductServiceTestSuite) TestCreate() {
	testCases := []struct {
		name    string
		product entity.Product
		err     entity.CustomError
	}{
		{
			name:    "Test case 1: success create product | 200",
			product: ts.product,
			err:     ts.customError,
		},
		{
			name:    "Test case 2: request is invalid | 400",
			product: entity.Product{},
			err: entity.CustomError{
				HttpCode: http.StatusBadRequest,
				Err:      errors.New(fmt.Sprintf(entity.ErrFieldInvalid, "brand_id")),
			},
		},
	}

	for _, tc := range testCases {
		ctx, cancel := context.WithTimeout(context.Background(), *dbDuration)

		ts.mockRepository.On("Create", ctx, tc.product).Return(tc.err).Once()
		err := ts.serviceInstance.Create(context.Background(), tc.product)

		ts.Equal(tc.err, err)

		cancel()
	}
}

func (ts *ProductServiceTestSuite) TestGet() {
	testCases := []struct {
		name    string
		id      int64
		product entity.Product
		err     entity.CustomError
	}{
		{
			name:    "Test case 1: success get product | 200",
			product: ts.product,
			id:      1,
			err:     ts.customError,
		},
		{
			name:    "Test case 2: record is not found | 404",
			product: entity.Product{},
			err: entity.CustomError{
				HttpCode: http.StatusNotFound,
				Err:      errors.New(entity.ErrorRecordNotExist),
			},
		},
	}

	for _, tc := range testCases {
		ctx, cancel := context.WithTimeout(context.Background(), *dbDuration)

		ts.mockRepository.On("Get", ctx, tc.id).Return(tc.product, tc.err).Once()
		product, err := ts.serviceInstance.Get(context.Background(), tc.id)

		ts.Equal(tc.product, product)
		ts.Equal(tc.err, err)

		cancel()
	}
}

func TestProductService(t *testing.T) {
	suite.Run(t, new(ProductServiceTestSuite))
}
