package balance

import (
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/db"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
)

type balanceRepo struct {
	db   db.DB
	psql squirrel.StatementBuilderType
}

// Inc implements BalanceRepository.
func (b balanceRepo) Inc(uID int, count float64) error {
	sql, args, err := b.psql.Update("users").
		Where("id = ?", uID).
		Set("balance", squirrel.Expr("balance + ?", count)).
		ToSql()

	if err != nil {
		return err
	}

	_, err = b.db.DB.Exec(sql, args...)

	return err
}

// Dec implements BalanceRepository.
func (b balanceRepo) Dec(uID int, count float64) error {
	sql, args, err := b.psql.Update("users").
		Where("id = ?", uID).
		Set("balance", squirrel.Expr("balance - ?", count)).
		Set("balance_wd", squirrel.Expr("balance_wd - ?", count*-1)).
		ToSql()

	if err != nil {
		return err
	}

	_, err = b.db.DB.Exec(sql, args...)

	return err
}

// Get implements BalanceRepository.
func (b balanceRepo) Get(uID int) (*models.UserBalance, error) {
	sql, args, err := b.psql.Select("balance", "balance_wd").
		From("users").
		Where("id = ?", uID).
		ToSql()

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	rows, err := b.db.DB.Query(sql, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	if rows.Next() {
		balance := models.UserBalance{}

		rows.Scan(&balance.Balance, &balance.BalanceWD)
		return &balance, nil
	}

	return nil, errors.New("no data")
}

func NewBalanceRepo(db db.DB) balanceRepository {
	return balanceRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
