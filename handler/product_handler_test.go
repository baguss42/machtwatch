package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/baguss42/machtwatch/entity"
	"github.com/baguss42/machtwatch/test/factory"
	mock "github.com/baguss42/machtwatch/test/mock/service"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ProductHandlerTestSuite struct {
	suite.Suite
	mockService     *mock.ProductServiceMock
	handlerInstance ProductHandler
	product           entity.Product
	customError     entity.CustomError
}

func (ts *ProductHandlerTestSuite) SetupTest() {
	ts.mockService = &mock.ProductServiceMock{}
	ts.handlerInstance = ProductHandler{Service: ts.mockService}
	ts.product = (&factory.ProductFactory{}).Build()
	ts.customError = entity.NewCustomError()
}

func (ts *ProductHandlerTestSuite) TestDecodeBrand() {
	productByte, _ := json.Marshal(ts.product)
	bodyRequest := bytes.NewReader(productByte)
	product, err := DecodeProduct(bodyRequest)
	ts.Equal(product, ts.product)
	ts.Nil(err)

	bodyRequest = bytes.NewReader([]byte("abc"))
	product, err = DecodeProduct(bodyRequest)
	ts.NotEqual(product, ts.product)
	ts.NotNil(err)
}

func (ts *ProductHandlerTestSuite) TestCreate() {
	testCases := []struct {
		name          string
		brand         entity.Product
		id int64
		method        string
		err           entity.CustomError
		isErrorDecode bool
		httpStatus    int
	}{
		{
			name:       "Test case 1: success create brand | 200",
			brand:      ts.product,
			id: 1,
			method:     http.MethodPost,
			err:        ts.customError,
			httpStatus: http.StatusCreated,
		},
		{
			name:          "Test case 2: failed decode brand | 400",
			brand:         entity.Product{},
			id: 0,
			method:        http.MethodPost,
			isErrorDecode: true,
			err:           entity.CustomError{HttpCode: http.StatusBadRequest, Err: errors.New("could not decode body")},
			httpStatus:    http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		brandByte, _ := json.Marshal(tc.brand)
		if tc.isErrorDecode {
			brandByte, _ = json.Marshal([]byte(""))
		}
		r := httptest.NewRequest(tc.method, "/product", bytes.NewReader(brandByte))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		ts.mockService.On("Create", r.Context(), tc.brand).Return(tc.err).Once()

		_, err := ts.handlerInstance.Create(w, r)
		res := w.Result()
		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			ts.Error(err, "unexpected error")
		}

		ts.Equal(tc.httpStatus, res.StatusCode)

		res.Body.Close()
	}
}

func TestProductHandler(t *testing.T) {
	suite.Run(t, new(ProductHandlerTestSuite))
}
