package entity

import "time"

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
