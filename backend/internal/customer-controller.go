package internal

import (
	"encoding/json"
	"net/http"
	"time"
)

type CustomerSchema struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	customer := CustomerSchema{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	repository := CustomerRepository{}
	customerEntity := Customer{
		ID:        len(repository.customers) + 1,
		Name:      customer.Name,
		Email:     customer.Email,
		Phone:     customer.Phone,
		Address:   customer.Address,
		CreatedAt: time.Now().String(),
	}
	repository.Create(customerEntity)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customerEntity)
}
