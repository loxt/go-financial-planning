package http

import (
	"github.com/loxt/go-financial-planning/adapter/http/transaction"
	"net/http"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transaction/create", transaction.CreateTransaction)

	_ = http.ListenAndServe(":8080", nil)
}
