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

type TransactionHandlerTestSuite struct {
	suite.Suite
	mockService        *mock.TransactionServiceMock
	handlerInstance    TransactionHandler
	transactionOrder   entity.TransactionOrder
	transactionDetails []entity.TransactionDetail
	customError        entity.CustomError
}

func (ts *TransactionHandlerTestSuite) SetupTest() {
	ts.mockService = &mock.TransactionServiceMock{}
	ts.handlerInstance = TransactionHandler{Service: ts.mockService}
	ts.transactionOrder = (&factory.TransactionOrderFactory{}).Build()
	ts.transactionDetails = (&factory.TransactionDetailsFactory{}).Build()
	ts.customError = entity.NewCustomError()
}

func (ts *TransactionHandlerTestSuite) TestDecodeTransactionOrder() {
	transactionOdrderByte, _ := json.Marshal(ts.transactionOrder)
	bodyRequest := bytes.NewReader(transactionOdrderByte)
	transactionOrder, err := DecodeTransactionOrder(bodyRequest)
	ts.Equal(transactionOrder, ts.transactionOrder)
	ts.Nil(err)

	bodyRequest = bytes.NewReader([]byte("abc"))
	transactionOrder, err = DecodeTransactionOrder(bodyRequest)
	ts.NotEqual(transactionOrder, ts.transactionOrder)
	ts.NotNil(err)
}

func (ts *TransactionHandlerTestSuite) TestCreate() {
	testCases := []struct {
		name             string
		transactionOrder entity.TransactionOrder
		id               int64
		err              entity.CustomError
		isErrorDecode    bool
		httpMethod       string
		httpStatus       int
	}{
		{
			name:             "Test case 1: success create brand | 201",
			transactionOrder: ts.transactionOrder,
			id:               1,
			err:              ts.customError,
			httpMethod:       http.MethodPost,
			httpStatus:       http.StatusCreated,
		},
		{
			name:             "Test case 2: failed decode brand | 400",
			transactionOrder: entity.TransactionOrder{},
			id:               0,
			isErrorDecode:    true,
			err:              entity.CustomError{HttpCode: http.StatusBadRequest, Err: errors.New("could not decode body")},
			httpMethod:       http.MethodPost,
			httpStatus:       http.StatusBadRequest,
		},
		{
			name:             "Test case 3: method not allowed | 405",
			transactionOrder: ts.transactionOrder,
			id:               1,
			err:              ts.customError,
			httpMethod:       http.MethodPut,
			httpStatus:       http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range testCases {
		brandByte, _ := json.Marshal(tc.transactionOrder)
		if tc.isErrorDecode {
			brandByte, _ = json.Marshal([]byte(""))
		}
		r := httptest.NewRequest(tc.httpMethod, "/order", bytes.NewReader(brandByte))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		ts.mockService.On("Create", r.Context(), tc.transactionOrder).Return(tc.err).Once()

		_, err := ts.handlerInstance.Transaction(w, r)
		res := w.Result()
		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			ts.Error(err, "unexpected error")
		}

		ts.Equal(tc.httpStatus, res.StatusCode)

		res.Body.Close()
	}
}

func (ts *TransactionHandlerTestSuite) TestGet() {
	testCases := []struct {
		name               string
		transactionDetails []entity.TransactionDetail
		id                 int64
		err                entity.CustomError
		httpMethod         string
		httpStatus         int
	}{
		{
			name:               "Test case 1: success get brand | 200",
			transactionDetails: ts.transactionDetails,
			id:                 1,
			err:                ts.customError,
			httpMethod:         http.MethodGet,
			httpStatus:         http.StatusOK,
		},
		{
			name:               "Test case 2: invalid id | 400",
			transactionDetails: []entity.TransactionDetail{},
			id:                 0,
			err:                entity.CustomError{HttpCode: http.StatusBadRequest, Err: errors.New("invalid id")},
			httpMethod:         http.MethodGet,
			httpStatus:         http.StatusBadRequest,
		},
		{
			name:               "Test case 3: record not found | 404",
			transactionDetails: []entity.TransactionDetail{},
			id:                 1001,
			err:                entity.CustomError{HttpCode: http.StatusNotFound, Err: errors.New("could not decode body")},
			httpMethod:         http.MethodGet,
			httpStatus:         http.StatusNotFound,
		},
		{
			name:               "Test case 4: method not allowed | 405",
			transactionDetails: ts.transactionDetails,
			id:                 1,
			err:                ts.customError,
			httpMethod:         http.MethodDelete,
			httpStatus:         http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range testCases {
		r := httptest.NewRequest(tc.httpMethod, "/order", nil)
		q := r.URL.Query()
		idStr := strconv.FormatInt(tc.id, 10)
		q.Add("id", idStr)
		r.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()

		ts.mockService.On("Get", r.Context(), tc.id).Return(tc.transactionDetails, tc.err).Once()

		_, err := ts.handlerInstance.Transaction(w, r)
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

func TestTransactionHandler(t *testing.T) {
	suite.Run(t, new(TransactionHandlerTestSuite))
}
