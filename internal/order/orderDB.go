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

// GetOrder implements orderRepository.
func (o orderRepo) GetOrder(number string) (int, error) {
	sql, args, err := o.psql.Select("uid").
		From("orders").
		Where("number = ?", number).
		ToSql()

	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	rows, err := o.db.DB.Query(sql, args...)

	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	if rows.Next() {
		var uId int

		err = rows.Scan(&uId)

		if err != nil {
			logrus.Error(err)
			return 0, err
		}

		return uId, nil
	}

	return 0, nil
}

// Add implements OrderRepository.
func (o orderRepo) Add(uID int, number string) error {
	currTime := time.Now().Format("2006-01-02 15:04:05")
	sql, args, err := o.psql.Insert("orders").
		Columns("uid", "number", "bonuses", "status", "created_at", "updated_at").
		Values(uID, number, 0, models.RegisteredStatus, currTime, currTime).
		ToSql()

	if err != nil {
		logrus.Error(err)
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
