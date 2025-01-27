package order

import (
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/db"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
)

type orderRepo struct {
	db   db.DB
	psql squirrel.StatementBuilderType
}

// Add implements OrderRepository.
func (o orderRepo) Add(uID int, order models.Order) error {
	currTime := time.Now().String()
	sql, args, err := o.psql.Insert("orders").Columns("uid", "number", "bonuses", "status", "created_at", "updated_at").
		Values(uID).Values(order.Number).Values(0).Values(order.Status).Values(currTime).Values(currTime).ToSql()

	if err != nil {
		return err
	}

	_, err = o.db.DB.Exec(sql, args...)

	return err
}

// Get implements OrderRepository.
func (o orderRepo) Get(uId int) ([]models.Order, error) {
	orders := make([]models.Order, 0)
	sql, args, err := o.psql.Select("id", "number", "bonuses", "status").
		From("orders").Where("uID = ?", uId).ToSql()

	if err != nil {
		logrus.Error(err)
		return orders, err
	}

	rows, err := o.db.DB.Query(sql, args...)

	if err != nil {
		logrus.Error(err)
		return orders, err
	}

	defer rows.Close()

	for rows.Next() {
		var order models.Order

		err = rows.Scan(&order.ID, &order.Number, &order.BonusesCount, &order.Status)

		if err != nil {
			logrus.Error(err)
			continue
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func NewOrderRepo(db db.DB) orderRepository {
	return orderRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
