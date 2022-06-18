package main

import (
	"fmt"
	"github.com/baguss42/machtwatch/database"
	"github.com/baguss42/machtwatch/handler"
	"github.com/baguss42/machtwatch/middleware"
	"github.com/baguss42/machtwatch/service"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	_ = godotenv.Load()
	middleware.Load()

	db := database.Load()
	service.InitDBTimeOut()

	brandHandler := handler.BrandHandler{Service: service.NewBrandService(db)}
	http.HandleFunc("/brand", middleware.Apply(brandHandler.Create))

	productHandler := handler.ProductHandler{Service: service.NewProductService(db)}
	http.HandleFunc("/product", middleware.Apply(productHandler.Product))

	transactionHandler := handler.TransactionHandler{Service: service.NewTransactionService(db)}
	http.HandleFunc("/order", middleware.Apply(transactionHandler.Transaction))

	fmt.Println("server starting at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}