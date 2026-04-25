# TxnFlow Engine

A simple payment transaction simulation system built with Go.

---

## 🚀 Features

- Create transaction
- Complete transaction using ISO 8583 response codes
- Transaction lifecycle:
  - PENDING → SUCCESS
  - PENDING → FAILED

---

## 📦 API Endpoints

### Create Transaction

POST /transactions

Request:
{
  "amount": 150000
}

Response:
{
  "id": "TXN-xxxx",
  "amount": 150000,
  "status": "PENDING",
  "iso_code": "",
  "message": "Awaiting processor result"
}

---

### Complete Transaction

POST /simulate/complete?id=TXN-xxxx

Request:
{
  "iso_code": "51"
}

Response:
{
  "id": "TXN-xxxx",
  "amount": 150000,
  "status": "FAILED",
  "iso_code": "51",
  "message": "Insufficient funds"
}

---

## 🔁 Transaction Flow

Create Transaction → PENDING  
                   ↓  
          Complete Transaction  
                   ↓  
        SUCCESS (00) / FAILED (others)

---

## 🧠 ISO 8583 Codes

00 → Approved  
14 → Invalid card  
51 → Insufficient funds  
54 → Expired card  
91 → Issuer unavailable  

---

## 🛠 Tech Stack

- Go (net/http)
- In-memory storage (map)

---

## ▶️ Run Locally

go run .

---

## 🎯 Purpose

This project simulates real-world payment transaction flows, including:

- Status transitions
- Failure handling
- ISO 8583 response mapping

Built to demonstrate system-level thinking for SDET / QA Automation roles.

---

## 🔥 Next Improvements

- Idempotency handling
- Retry logic
- Playwright API testing
- Persistent storage (SQLite)
