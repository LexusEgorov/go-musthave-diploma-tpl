package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
	servererrors "github.com/LexusEgorov/go-musthave-diploma-tpl/internal/serverErrors"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/utils"
	"github.com/sirupsen/logrus"
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
	Add(uID int, number string) error
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
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(body, &userData); err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.user.Register(userData)

	if err != nil {
		if errors.As(err, &servererrors.ConflictError{}) {
			w.WriteHeader(http.StatusConflict)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", user.Jwt)
}

func (h handlers) Auth(w http.ResponseWriter, r *http.Request) {
	var userData models.User
	body, err := io.ReadAll(r.Body)

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(body, &userData); err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.user.Auth(userData)

	if err != nil {
		if errors.As(err, &servererrors.UnauthorizedError{}) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", user.Jwt)
}

func (h handlers) AddOrders(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jwt := r.Header.Get("Authorization")

	uId, err := utils.ValidateJWT(jwt)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.order.Add(uId, string(body))

	if err != nil {
		if errors.As(err, &servererrors.OkayError{}) {
			w.WriteHeader(http.StatusOK)
			return
		}

		if errors.As(err, &servererrors.ConflictError{}) {
			w.WriteHeader(http.StatusConflict)
			return
		}

		if errors.As(err, &servererrors.WrongNumberError{}) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h handlers) GetOrders(w http.ResponseWriter, r *http.Request) {}

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
