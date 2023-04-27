package server

import (
	"encoding/json"
	"net/http"

	"github.com/gabrielluizsf/paymentQueue/payment"
)

func Start() {
	paymentQueue := payment.NewQueue()
  
	http.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
	  if r.Method != "POST" {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	  }
  
	  var payment payment.Payment
	  err := json.NewDecoder(r.Body).Decode(&payment)
	  if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	  }
  
	  paymentQueue.AddPayment(payment)
	  w.WriteHeader(http.StatusCreated)
	})
  
	http.HandleFunc("/process_payment", func(w http.ResponseWriter, r *http.Request) {
	  if r.Method != "GET" {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	  }
  
	  payment := paymentQueue.ProcessPayment()
	  if payment == nil {
		http.Error(w, "No payments to process", http.StatusNotFound)
		return
	  }
  
	  json.NewEncoder(w).Encode(payment)
	})
  
	http.ListenAndServe(":8080", nil)
  }
  