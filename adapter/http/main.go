package http

import (
	"github.com/loxt/go-financial-planning/adapter/http/actuator"
	"github.com/loxt/go-financial-planning/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// Init listen the server
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transaction/create", transaction.CreateTransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	_ = http.ListenAndServe(":8080", nil)
}
