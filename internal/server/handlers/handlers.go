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
	Inc(uID int, count float64) error
	Dec(uID int, count float64) error
	Get(uID int) *models.UserBalance
	IsEnough(uID int, count float64) bool
}

type UserManager interface {
	Register(u models.User) (*models.UserAuth, error)
	Auth(u models.User) (*models.UserAuth, error)
}

type OrderManager interface {
	Add(uID int, number string, wdSum *float64) error
	Get(uID int) []models.Order
	GetWithdrawals(uID int) []models.Withdrawal
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

	uID, _ := utils.ValidateJWT(jwt)

	err = h.order.Add(uID, string(body), nil)

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

func (h handlers) GetOrders(w http.ResponseWriter, r *http.Request) {
	jwt := r.Header.Get("Authorization")

	uID, _ := utils.ValidateJWT(jwt)

	orders := h.order.Get(uID)

	if len(orders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response, err := json.Marshal(orders)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h handlers) GetBalance(w http.ResponseWriter, r *http.Request) {
	jwt := r.Header.Get("Authorization")

	uID, _ := utils.ValidateJWT(jwt)

	balance := h.balance.Get(uID)

	response, err := json.Marshal(balance)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h handlers) WithdrawBalance(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jwt := r.Header.Get("Authorization")

	uID, _ := utils.ValidateJWT(jwt)
	request := models.WdBalance{}

	if err = json.Unmarshal(body, &request); err != nil {
		var syntaxError *json.SyntaxError
		if errors.As(err, &syntaxError) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !h.balance.IsEnough(uID, request.Sum) {
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	err = h.order.Add(uID, request.Order, &request.Sum)

	if err != nil {
		if errors.As(err, &servererrors.WrongNumberError{}) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.balance.Dec(uID, request.Sum)

	if err != nil {
		logrus.Error(err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h handlers) GetWithdrawals(w http.ResponseWriter, r *http.Request) {
	jwt := r.Header.Get("Authorization")

	uID, _ := utils.ValidateJWT(jwt)

	withdrawals := h.order.GetWithdrawals(uID)

	if len(withdrawals) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response, err := json.Marshal(withdrawals)

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func NewHandlers(user UserManager, balance BalanceManager, order OrderManager) handlers {
	return handlers{
		user:    user,
		balance: balance,
		order:   order,
	}
}
