package main

import (
	"desafiogolang-payment/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define os endpoints
	r.HandleFunc("/process-payment", handlers.ProcessPayment).Methods("POST")
	r.HandleFunc("/payment-status", handlers.GetPaymentStatus).Methods("GET")
	r.HandleFunc("/convert-currency", handlers.ConvertCurrency).Methods("POST")

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
