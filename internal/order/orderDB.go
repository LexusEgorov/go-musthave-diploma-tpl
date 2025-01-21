package order

import (
	"github.com/Masterminds/squirrel"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/db"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
)

type orderRepo struct {
	db   db.DB
	psql squirrel.StatementBuilderType
}

// Add implements OrderRepository.
func (orderRepo) Add(uID int, o models.Order) error {
	panic("unimplemented")
}

// Get implements OrderRepository.
func (o orderRepo) Get(uId int) []models.Order {
	panic("unimplemented")
}

func NewOrderRepo(db db.DB) orderRepository {
	return orderRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
