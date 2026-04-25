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

var isoMap = map[string]string{
	"00": "Approved",
	"14": "Invalid card",
	"51": "Insufficient funds",
	"54": "Expired card",
	"91": "Issuer unavailable",
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

func completeTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	tx, exists := store[id]
	if !exists {
		http.Error(w, "transaction not found", http.StatusNotFound)
		return
	}

	var req struct {
		ISOCode string `json:"iso_code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	message, ok := isoMap[req.ISOCode]
	if !ok {
		http.Error(w, "unknown iso code", http.StatusBadRequest)
		return
	}

	if req.ISOCode == "00" {
		tx.Status = "SUCCESS"
	} else {
		tx.Status = "FAILED"
	}

	tx.ISOCode = req.ISOCode
	tx.Message = message

	store[id] = tx

	w.Header().Set("Content Type", "application/json")
	json.NewEncoder(w).Encode(tx)
}

func main() {
	http.HandleFunc("/transactions", createTransactionHandler)
	http.HandleFunc("/simulate/complete", completeTransactionHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("server running on :8080")
	http.ListenAndServe(":8080", nil)
}
