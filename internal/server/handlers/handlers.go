package handlers

import (
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
	Register(u models.User) (models.UserAuth, error)
	Auth(u models.User) (models.UserAuth, error)
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

func (h handlers) Registration(w http.ResponseWriter, r *http.Request) {}

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
