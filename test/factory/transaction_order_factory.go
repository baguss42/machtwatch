package factory

import (
	"github.com/baguss42/machtwatch/entity"
)

type TransactionOrderFactory struct {
	entity.Product
}

func (f *TransactionOrderFactory) Build() entity.TransactionOrder {
	return entity.TransactionOrder{
		Carts: []entity.Carts{
			{
				ProductID: 1,
				Quantity:  5,
			},
			{
				ProductID: 2,
				Quantity:  5,
			},
		},
	}
}
