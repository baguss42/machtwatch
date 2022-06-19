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

type BrandServiceTestSuite struct {
	suite.Suite
	mockRepository  *mock.BrandRepositoryMock
	serviceInstance BrandService
	brand           entity.Brand
	customError     entity.CustomError
}

func (ts *BrandServiceTestSuite) SetupTest() {
	ts.mockRepository = &mock.BrandRepositoryMock{}
	ts.serviceInstance = BrandService{Repository: ts.mockRepository}
	ts.brand = (&factory.BrandFactory{}).Build()
	ts.customError = entity.NewCustomError()
	InitDBTimeOut()
}

func (ts *BrandServiceTestSuite) TestNewBrandService() {
	var db *sql.DB
	brandService := NewBrandService(db)
	ts.Equal(&BrandService{Repository: &repository.BrandRepository{DB: db}}, brandService)
}

func (ts *BrandServiceTestSuite) TestCreate() {
	testCases := []struct {
		name  string
		brand entity.Brand
		err   entity.CustomError
	}{
		{
			name:  "Test case 1: success create brand | 200",
			brand: ts.brand,
			err:   ts.customError,
		},
		{
			name:  "Test case 2: request is invalid | 400",
			brand: entity.Brand{},
			err: entity.CustomError{
				HttpCode: http.StatusBadRequest,
				Err:      errors.New(fmt.Sprintf(entity.ErrFieldRequired, "name")),
			},
		},
	}

	for _, tc := range testCases {
		ctx, cancel := context.WithTimeout(context.Background(), *dbDuration)

		ts.mockRepository.On("Create", ctx, tc.brand).Return(tc.err).Once()
		err := ts.serviceInstance.Create(context.Background(), tc.brand)

		ts.Equal(tc.err, err)

		cancel()
	}
}

func TestBrandService(t *testing.T) {
	suite.Run(t, new(BrandServiceTestSuite))
}
