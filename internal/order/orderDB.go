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

// Update implements orderRepository.
func (o orderRepo) Update(order models.AccuralOrder) (*models.UserUpdate, error) {
	sql, args, err := o.psql.Update("orders").
		Set("status", order.Status).
		Set("bonuses", order.Count).
		Where("number = ?", order.ID).
		Suffix("RETURNING bonuses, uId").
		ToSql()

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	update := models.UserUpdate{}
	err = o.db.DB.QueryRow(sql, args...).Scan(&update.Count, &update.ID)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &update, nil
}

// GetQueue implements orderRepository.
func (o orderRepo) GetQueue() []string {
	queue := make([]string, 0)
	sql, args, err := o.psql.Select("number").
		From("orders").
		Where("status != ?", models.ProcessedStatus).
		Where("status != ?", models.InvalidStatus).
		ToSql()

	if err != nil {
		logrus.Error(err)
		return queue
	}

	rows, err := o.db.DB.Query(sql, args...)

	if err != nil {
		logrus.Error(err)
		return queue
	}

	for rows.Next() {
		var row string

		err = rows.Scan(&row)

		if err != nil {
			logrus.Error(err)
			break
		}

		queue = append(queue, row)
	}

	return queue
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
		var uID int

		err = rows.Scan(&uID)

		if err != nil {
			logrus.Error(err)
			return 0, err
		}

		return uID, nil
	}

	return 0, nil
}

// Add implements OrderRepository.
func (o orderRepo) Add(uID int, number string, count float64) error {
	status := models.RegisteredStatus

	if count != 0 {
		status = models.ProcessedStatus
	}

	currTime := time.Now().Format("2006-01-02 15:04:05")
	sql, args, err := o.psql.Insert("orders").
		Columns("uid", "number", "bonuses", "status", "created_at", "updated_at").
		Values(uID, number, count, status, currTime, currTime).
		ToSql()

	if err != nil {
		logrus.Error(err)
		return err
	}

	_, err = o.db.DB.Exec(sql, args...)

	return err
}

// Get implements OrderRepository.
func (o orderRepo) Get(uID int) ([]models.Order, error) {
	orders := make([]models.Order, 0)
	sql, args, err := o.psql.Select("number", "bonuses", "status", "created_at").
		From("orders").
		Where("uID = ?", uID).
		OrderBy("created_at").
		ToSql()

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

		err = rows.Scan(&order.Number, &order.BonusesCount, &order.Status, &order.CreatedAt)

		if err != nil {
			logrus.Error(err)
			continue
		}

		orders = append(orders, order)
	}

	return orders, nil
}

// GetWithdrawals implements OrderRepository.
func (o orderRepo) GetWithdrawals(uID int) ([]models.Withdrawal, error) {
	withdrawals := make([]models.Withdrawal, 0)
	sql, args, err := o.psql.Select("number", "bonuses", "created_at").
		From("orders").
		Where("uid = ?", uID).
		Where("bonuses < 0").
		OrderBy("created_at desc").
		ToSql()

	if err != nil {
		logrus.Error(err)
		return withdrawals, err
	}

	rows, err := o.db.DB.Query(sql, args...)

	if err != nil {
		logrus.Error(err)
		return withdrawals, err
	}

	defer rows.Close()

	for rows.Next() {
		var w models.Withdrawal

		err = rows.Scan(&w.Order, &w.Sum, &w.ProcessedAt)

		if err != nil {
			logrus.Error(err)
			continue
		}

		w.Sum = w.Sum * -1

		withdrawals = append(withdrawals, w)
	}

	return withdrawals, nil
}

func NewOrderRepo(db db.DB) orderRepository {
	return orderRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
