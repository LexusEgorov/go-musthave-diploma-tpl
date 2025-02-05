package order

import (
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
	servererrors "github.com/LexusEgorov/go-musthave-diploma-tpl/internal/serverErrors"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/utils"
	"github.com/sirupsen/logrus"
)

type orderRepository interface {
	Add(uID int, o string) error
	Get(uId int) ([]models.Order, error)
	GetOrder(o string) (int, error)
}

type order struct {
	repo orderRepository
}

//TODO: resolve errors with handlers

// Add implements handlers.OrderManager.
func (o order) Add(uID int, number string) error {
	if !utils.LunaCheck(number) {
		return servererrors.WrongNumberError{Number: number}
	}

	user, err := o.repo.GetOrder(number)

	if err != nil {
		logrus.Error(err)
		return err
	}

	if user == 0 {
		return o.repo.Add(uID, number)
	}

	if user == uID {
		return servererrors.OkayError{}
	}

	return servererrors.ConflictError{Used: number}
}

// Get implements handlers.OrderManager.
func (o order) Get(uID int) []models.Order {
	order, err := o.repo.Get(uID)

	if err != nil {
		logrus.Error(err)
		return nil
	}

	return order
}

func NewOrder(repo orderRepository) order {
	return order{
		repo: repo,
	}
}
