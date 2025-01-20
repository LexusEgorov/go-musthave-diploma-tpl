package order

import "github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"

type orderRepo struct {
}

// Add implements OrderRepository.
func (orderRepo) Add(uID int, o models.Order) error {
	panic("unimplemented")
}

// Get implements OrderRepository.
func (o orderRepo) Get(uId int) []models.Order {
	panic("unimplemented")
}

func NewOrderRepo(connection string) orderRepository {
	return orderRepo{}
}
