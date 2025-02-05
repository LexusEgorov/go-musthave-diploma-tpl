package balance

import (
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
	"github.com/sirupsen/logrus"
)

type balanceRepository interface {
	Inc(uID int, count float64) error
	Dec(uID int, count float64) error
	Get(uID int) (*models.UserBalance, error)
}

type balance struct {
	repo balanceRepository
}

// IsEnough implements handlers.BalanceManager.
func (b balance) IsEnough(uID int, count float64) bool {
	currBalance, err := b.repo.Get(uID)

	if err != nil {
		logrus.Error(err)
		return false
	}

	return currBalance.Balance >= count
}

// Dec implements handlers.BalanceManager.
func (b balance) Dec(uID int, count float64) error {
	return b.repo.Dec(uID, count)
}

// Inc implements handlers.BalanceManager.
func (b balance) Inc(uID int, count float64) error {
	return b.repo.Inc(uID, count)
}

// Get implements handlers.BalanceManager.
func (b balance) Get(uID int) *models.UserBalance {
	balance, err := b.repo.Get(uID)

	if err != nil {
		logrus.Error(err)
		return nil
	}

	return balance
}

func NewBalance(repo balanceRepository) balance {
	return balance{
		repo: repo,
	}
}
