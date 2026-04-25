package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Transaction struct {
	ID      string `json:"id"`
	Amount  int    `json:"amount"`
	Status  string `json:"status"`
	ISOCode string `json:"iso_code"`
	Message string `json:"message"`
}

var store = map[string]Transaction{}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func createTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Amount int `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Amount <= 0 {
		http.Error(w, "amount must be greater than zero", http.StatusBadRequest)
		return
	}

	tx := Transaction{
		ID:      fmt.Sprintf("TXN-%d", time.Now().UnixNano()),
		Amount:  req.Amount,
		Status:  "PENDING",
		ISOCode: "",
		Message: "Awaiting processor result",
	}

	store[tx.ID] = tx

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tx)
}

func main() {
	http.HandleFunc("/transactions", createTransactionHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("server running on :8080")
	http.ListenAndServe(":8080", nil)
}
