package entity

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

const (
	errorRecordNotExist = "record not exist"
	errorUnexpected     = "unexpected error"
	errorUnauthorized   = "unauthorized error"
	errorForbidden      = "forbidden error"
)

var (
	ErrFieldRequired = "field %s is required"
	ErrFieldInvalid  = "field %s is invalid"
)

type CustomError struct {
	HttpCode int
	Err      error
}

func NewCustomError() CustomError {
	return CustomError{
		HttpCode: http.StatusOK, // By default status 200
		Err:      nil,
	}
}

func (c *CustomError) ErrorNotFound() {
	c.HttpCode = http.StatusNotFound
	c.Err = errors.New(errorRecordNotExist)
}

func (c *CustomError) ErrorUnexpected(err error) {
	c.HttpCode = http.StatusInternalServerError
	if err != nil {
		c.Err = err
	}
}

func (c *CustomError) ErrorBadRequest(err error) {
	c.HttpCode = http.StatusBadRequest
	if err != nil {
		c.Err = err
	}
}

func (c *CustomError) ErrorRecordAlreadyExist(err error) {
	c.HttpCode = http.StatusConflict
	if err != nil {
		c.Err = err
	}
}

func (c *CustomError) ErrorUnprocessableEntity(err error) {
	c.HttpCode = http.StatusUnprocessableEntity
	if err != nil {
		c.Err = err
	}
}

func (c *CustomError) BuildSQLError(method string) {
	httpCode := map[string]int{
		"create": 201,
		"get":    200,
		"update": 200,
	}
	if c.Err != nil {
		if mysqlError, ok := c.Err.(*mysql.MySQLError); ok {
			switch {
			case mysqlError.Number == 1062: // duplicate record
				c.ErrorRecordAlreadyExist(errors.New("record already exist"))
			case mysqlError.Number == 1452: //  foreign key constraint fails
				c.ErrorUnprocessableEntity(errors.New("product brand is not found"))
			default:
				c.ErrorUnexpected(nil)
			}
		} else if c.Err == sql.ErrNoRows {
			c.ErrorNotFound()
		}
	} else {
		c.HttpCode = httpCode[method]
	}
}
