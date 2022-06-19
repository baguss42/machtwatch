package factory

import (
	"github.com/baguss42/machtwatch/entity"
)

type ProductFactory struct {
	entity.Product
}

func (f *ProductFactory) Build() entity.Product {
	return entity.Product{
		ID: 1,
		Title:        "Seiko Watch",
		Description: "Seiko is one of the few fully integrated watch manufactures",
		Price:		 100000,
		PriceReduction:  25000,
		Stock: 100,
		IsActive: true,
	}
}
