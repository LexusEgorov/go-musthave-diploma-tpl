package balance

import (
	"github.com/Masterminds/squirrel"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/db"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
)

type balanceRepo struct {
	db   db.DB
	psql squirrel.StatementBuilderType
}

// Inc implements BalanceRepository.
func (b balanceRepo) Inc(uID int, count int) error {
	sql, args, err := b.psql.Update("users").Where("id = ?", uID).Set("balance", squirrel.Expr("balance + ?", count)).ToSql()

	if err != nil {
		return err
	}

	_, err = b.db.DB.Exec(sql, args...)

	return err
}

// Dec implements BalanceRepository.
func (b balanceRepo) Dec(uID int, count int) error {
	sql, args, err := b.psql.Update("users").Where("id = ?", uID).Set("balance", squirrel.Expr("balance - ?", count)).ToSql()

	if err != nil {
		return err
	}

	_, err = b.db.DB.Exec(sql, args...)

	return err
}

// Get implements BalanceRepository.
func (b balanceRepo) Get(uID int) int {
	balance := 0
	sql, args, err := b.psql.Select("balance").From("users").Where("id = ?", uID).ToSql()

	if err != nil {
		//TODO: LOGGING
		return balance
	}

	err = b.db.DB.QueryRow(sql, args...).Scan(&balance)

	if err != nil {
		//TODO: LOGGING
		return balance
	}

	return balance
}

// GetWithdrawals implements BalanceRepository.
func (b balanceRepo) GetWithdrawals(uID int) []models.Withdrawal {
	withdrawals := make([]models.Withdrawal, 0)
	sql, args, err := b.psql.Select("number", "bonuses", "created_at").From("orders").Where("uid = ?", uID).Where("bonuses < 0").ToSql()

	if err != nil {
		//TODO: LOGGING
		return withdrawals
	}

	rows, err := b.db.DB.Query(sql, args...)

	if err != nil {
		//TODO: LOGGING
		return withdrawals
	}

	defer rows.Close()

	for rows.Next() {
		var w models.Withdrawal

		err = rows.Scan(&w.Order, &w.Sum, &w.ProcessedAt)

		if err != nil {
			//TODO: LOGGING
			continue
		}

		w.Sum = w.Sum * -1

		withdrawals = append(withdrawals, w)
	}

	return withdrawals
}

func NewBalanceRepo(db db.DB) balanceRepository {
	return balanceRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
