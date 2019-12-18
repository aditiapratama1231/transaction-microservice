package main

import (
	"github.com/gin-gonic/gin"

	"transaction/config"
	"transaction/transaction"
)

func main() {
	r := gin.Default()
	db := config.DBInit()
	transaction := transaction.Transaction{DB: db}

	r.GET("/transactions", transaction.GetTransactions)
	r.POST("/transactions", transaction.CreateTransactions)

	r.Run()
}
