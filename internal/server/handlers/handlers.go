package handlers

import "net/http"

type handlers struct{}

func (h handlers) Registration(w http.ResponseWriter, r *http.Request) {}
func (h handlers) Auth(w http.ResponseWriter, r *http.Request)         {}

func (h handlers) GetOrders(w http.ResponseWriter, r *http.Request) {}
func (h handlers) AddOrders(w http.ResponseWriter, r *http.Request) {}

func (h handlers) GetBalance(w http.ResponseWriter, r *http.Request)      {}
func (h handlers) WithdrawBalance(w http.ResponseWriter, r *http.Request) {}

func (h handlers) GetWithdrawals(w http.ResponseWriter, r *http.Request) {}

func NewHandlers() handlers {
	return handlers{}
}
