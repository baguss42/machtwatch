package entity

import (
	"errors"
	"fmt"
	"time"
)

type Product struct {
	ID int64 `json:"id"`
	BrandID int64 `json:"brand_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price int64 `json:"price"`
	PriceReduction int64 `json:"price_reduction"`
	Stock int `json:"stock"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validate() error {
	if p.BrandID < 1 {
		return errors.New(fmt.Sprintf(ErrFieldInvalid, "brand_id"))
	}
	if p.Title == "" {
		return errors.New(fmt.Sprintf(ErrFieldRequired, "title"))
	}
	if p.Description == "" {
		return errors.New(fmt.Sprintf(ErrFieldRequired, "description"))
	}
	if p.Price < 1 {
		return errors.New(fmt.Sprintf(ErrFieldInvalid, "price"))
	}
	if p.PriceReduction < 0 {
		return errors.New(fmt.Sprintf(ErrFieldInvalid, "price_reduction"))
	}
	if p.Stock < 0 {
		return errors.New(fmt.Sprintf(ErrFieldInvalid, "stock"))
	}
	return nil
}