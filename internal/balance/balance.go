package balance

import (
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
	"github.com/sirupsen/logrus"
)

type balanceRepository interface {
	Inc(uID int, count int) error
	Dec(uID int, count int) error
	Get(uID int) (int, error)
	GetWithdrawals(uId int) ([]models.Withdrawal, error)
}

type balance struct {
	repo balanceRepository
}

//TODO: resolve errors with handlers

// Dec implements handlers.BalanceManager.
func (b balance) Dec(uID int, count int) int {
	b.repo.Dec(uID, count)
	return b.Get(uID)
}

// Get implements handlers.BalanceManager.
func (b balance) Get(uID int) int {
	balance, err := b.repo.Get(uID)

	if err != nil {
		logrus.Error(err)
		return 0
	}

	return balance
}

// GetWithdrawals implements handlers.BalanceManager.
func (b balance) GetWithdrawals(uID int) []models.Withdrawal {
	withdrawals, err := b.repo.GetWithdrawals(uID)

	if err != nil {
		logrus.Error(err)
	}

	return withdrawals
}

// Inc implements handlers.BalanceManager.
func (b balance) Inc(uID int, count int) int {
	b.repo.Inc(uID, count)
	return b.Get(uID)
}

func NewBalance(repo balanceRepository) balance {
	return balance{
		repo: repo,
	}
}
