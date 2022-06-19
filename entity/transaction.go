package entity

import "time"

type Transaction struct {
	ID        int64     `json:"id"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TransactionDetail struct {
	ID             int64     `json:"id"`
	TransactionID  int64     `json:"transaction_id"`
	ProductID      int64     `json:"product_id"`
	Price          int64     `json:"price"`
	PriceReduction int64     `json:"price_reduction"`
	Quantity       int       `json:"quantity"`
	FinalPrice     int64     `json:"final_price"`
	CreatedAt      time.Time `json:"created_at"`
}

type TransactionOrder struct { // Assume we create order form carts
	Carts []Carts `json:"carts"`
}

type Carts struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}
