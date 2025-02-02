package balance

import "github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"

type balanceRepository interface {
	Inc(uID int, count int) error
	Dec(uID int, count int) error
	Get(uID int) (int, error)
	GetWithdrawals(uId int) ([]models.Withdrawal, error)
}

type balance struct {
	repo balanceRepository
}

// Dec implements handlers.BalanceManager.
func (b balance) Dec(uID int, count int) (currBalance int, err error) {
	b.repo.Dec(uID, count)
	return b.repo.Get(uID)
}

// Get implements handlers.BalanceManager.
func (b balance) Get(uID int) (currBalance int, err error) {
	return b.repo.Get(uID)
}

// GetWithdrawals implements handlers.BalanceManager.
func (b balance) GetWithdrawals(uID int) ([]models.Withdrawal, error) {
	return b.repo.GetWithdrawals(uID)
}

// Inc implements handlers.BalanceManager.
func (b balance) Inc(uID int, count int) (currBalance int, err error) {
	b.repo.Inc(uID, count)
	return b.repo.Get(uID)
}

func NewBalance(repo balanceRepository) balance {
	return balance{
		repo: repo,
	}
}
