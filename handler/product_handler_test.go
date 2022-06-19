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
	"strconv"
	"testing"
)

type ProductHandlerTestSuite struct {
	suite.Suite
	mockService     *mock.ProductServiceMock
	handlerInstance ProductHandler
	product         entity.Product
	customError     entity.CustomError
}

func (ts *ProductHandlerTestSuite) SetupTest() {
	ts.mockService = &mock.ProductServiceMock{}
	ts.handlerInstance = ProductHandler{Service: ts.mockService}
	ts.product = (&factory.ProductFactory{}).Build()
	ts.customError = entity.NewCustomError()
}

func (ts *ProductHandlerTestSuite) TestDecodeProduct() {
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
		product       entity.Product
		id            int64
		err           entity.CustomError
		isErrorDecode bool
		httpMethod    string
		httpStatus    int
	}{
		{
			name:       "Test case 1: success create product | 201",
			product:    ts.product,
			id:         1,
			err:        ts.customError,
			httpMethod: http.MethodPost,
			httpStatus: http.StatusCreated,
		},
		{
			name:          "Test case 2: failed decode product | 400",
			product:       entity.Product{},
			id:            0,
			isErrorDecode: true,
			err:           entity.CustomError{HttpCode: http.StatusBadRequest, Err: errors.New("could not decode body")},
			httpMethod:    http.MethodPost,
			httpStatus:    http.StatusBadRequest,
		},
		{
			name:       "Test case 3: method not allowed | 405",
			product:    ts.product,
			id:         1,
			err:        ts.customError,
			httpMethod: http.MethodPut,
			httpStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range testCases {
		productByte, _ := json.Marshal(tc.product)
		if tc.isErrorDecode {
			productByte, _ = json.Marshal([]byte(""))
		}
		r := httptest.NewRequest(tc.httpMethod, "/product", bytes.NewReader(productByte))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		ts.mockService.On("Create", r.Context(), tc.product).Return(tc.err).Once()

		_, err := ts.handlerInstance.Product(w, r)
		res := w.Result()
		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			ts.Error(err, "unexpected error")
		}

		ts.Equal(tc.httpStatus, res.StatusCode)

		res.Body.Close()
	}
}

func (ts *ProductHandlerTestSuite) TestGet() {
	testCases := []struct {
		name       string
		product    entity.Product
		id         int64
		err        entity.CustomError
		httpMethod string
		httpStatus int
	}{
		{
			name:       "Test case 1: success get product | 200",
			product:    ts.product,
			id:         1,
			err:        ts.customError,
			httpMethod: http.MethodGet,
			httpStatus: http.StatusOK,
		},
		{
			name:       "Test case 2: invalid id | 400",
			product:    entity.Product{},
			id:         0,
			err:        entity.CustomError{HttpCode: http.StatusBadRequest, Err: errors.New("invalid id")},
			httpMethod: http.MethodGet,
			httpStatus: http.StatusBadRequest,
		},
		{
			name:       "Test case 3: record not found | 404",
			product:    entity.Product{},
			id:         1001,
			err:        entity.CustomError{HttpCode: http.StatusNotFound, Err: errors.New("could not decode body")},
			httpMethod: http.MethodGet,
			httpStatus: http.StatusNotFound,
		},
		{
			name:       "Test case 4: method not allowed | 405",
			product:    ts.product,
			id:         1,
			err:        ts.customError,
			httpMethod: http.MethodDelete,
			httpStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range testCases {
		r := httptest.NewRequest(tc.httpMethod, "/product", nil)
		q := r.URL.Query()
		idStr := strconv.FormatInt(tc.id, 10)
		q.Add("id", idStr)
		r.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()

		ts.mockService.On("Get", r.Context(), tc.id).Return(tc.product, tc.err).Once()

		_, err := ts.handlerInstance.Product(w, r)
		res := w.Result()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			ts.Error(err, "unexpected error")
		}

		ts.NotNil(data)
		ts.Equal(tc.httpStatus, res.StatusCode)

		res.Body.Close()
	}
}

func TestProductHandler(t *testing.T) {
	suite.Run(t, new(ProductHandlerTestSuite))
}
