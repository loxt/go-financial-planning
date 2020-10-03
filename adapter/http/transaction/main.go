package transaction

import (
	"encoding/json"
	"github.com/loxt/go-financial-planning/model/transaction"
	"github.com/loxt/go-financial-planning/util"
	"io/ioutil"
	"net/http"
)

// GetTransactions return all the transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salary",
			Amount:    1500.0,
			Type:      0,
			CreatedAt: util.StringToTime("2020-10-03T14:04:05"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

// CreateTransaction create a transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	_ = json.NewEncoder(w).Encode(res)
}
