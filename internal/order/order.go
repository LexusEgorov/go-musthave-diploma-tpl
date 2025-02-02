package order

import (
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
	"github.com/sirupsen/logrus"
)

type orderRepository interface {
	Add(uID int, o models.Order) error
	Get(uId int) ([]models.Order, error)
}

type order struct {
	repo orderRepository
}

//TODO: resolve errors with handlers

// Add implements handlers.OrderManager.
func (o order) Add(uID int, order models.Order) error {
	return o.repo.Add(uID, order)
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
