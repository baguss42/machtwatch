package factory

import (
	"github.com/baguss42/machtwatch/entity"
)

type TransactionDetailsFactory struct {
	entity.Product
}

func (f *TransactionDetailsFactory) Build() []entity.TransactionDetail {
	return []entity.TransactionDetail{
		{
			ID:             1,
			TransactionID:  1,
			ProductID:      1,
			Price:          50000,
			PriceReduction: 25000,
			Quantity:       2,
			FinalPrice:     50000,
		},
	}
}
