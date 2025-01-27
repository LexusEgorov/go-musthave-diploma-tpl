package order

import "github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"

type orderRepository interface {
	Add(uID int, o models.Order) error
	Get(uId int) ([]models.Order, error)
}

type order struct {
	repo orderRepository
}

// Add implements handlers.OrderManager.
func (order) Add(uID int, o models.Order) error {
	panic("unimplemented")
}

// Get implements handlers.OrderManager.
func (o order) Get(uID int) []models.Order {
	panic("unimplemented")
}

func NewOrder(repo orderRepository) order {
	return order{
		repo: repo,
	}
}
