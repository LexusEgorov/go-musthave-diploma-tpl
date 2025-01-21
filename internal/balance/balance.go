package balance

import "github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"

type balanceRepository interface {
	Inc(uID int, count int) error
	Dec(uID int, count int) error
	Get(uID int) int
	GetWithdrawals(uId int) []models.Withdrawal
}

type balance struct {
	repo balanceRepository
}

// Dec implements handlers.BalanceManager.
func (b balance) Dec(uID int, count int) (currBalance int) {
	panic("unimplemented")
}

// Get implements handlers.BalanceManager.
func (b balance) Get(uID int) int {
	panic("unimplemented")
}

// GetWithdrawals implements handlers.BalanceManager.
func (b balance) GetWithdrawals(uID int) []models.Withdrawal {
	panic("unimplemented")
}

// Inc implements handlers.BalanceManager.
func (b balance) Inc(uID int, count int) (currBalance int) {
	panic("unimplemented")
}

func NewBalance(repo balanceRepository) balance {
	return balance{
		repo: repo,
	}
}
