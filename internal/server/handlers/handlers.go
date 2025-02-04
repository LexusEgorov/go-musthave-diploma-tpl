package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
)

type BalanceManager interface {
	Inc(uID int, count int) (currBalance int)
	Dec(uID int, count int) (currBalance int)
	Get(uID int) (currBalance int)
	GetWithdrawals(uID int) []models.Withdrawal
}

type UserManager interface {
	Register(u models.User) (*models.UserAuth, error)
	Auth(u models.User) (*models.UserAuth, error)
}

type OrderManager interface {
	Add(uID int, o models.Order) error
	Get(uID int) []models.Order
}

type handlers struct {
	user    UserManager
	balance BalanceManager
	order   OrderManager
}

func (h handlers) Registration(w http.ResponseWriter, r *http.Request) {
	var userData models.User
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, &userData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.user.Register(userData)

	//TODO ERROR STATUS MAPPER
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h handlers) Auth(w http.ResponseWriter, r *http.Request) {}

func (h handlers) GetOrders(w http.ResponseWriter, r *http.Request) {}
func (h handlers) AddOrders(w http.ResponseWriter, r *http.Request) {}

func (h handlers) GetBalance(w http.ResponseWriter, r *http.Request)      {}
func (h handlers) WithdrawBalance(w http.ResponseWriter, r *http.Request) {}

func (h handlers) GetWithdrawals(w http.ResponseWriter, r *http.Request) {}

func NewHandlers(user UserManager, balance BalanceManager, order OrderManager) handlers {
	return handlers{
		user:    user,
		balance: balance,
		order:   order,
	}
}
