package balance

import "github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"

type balanceRepo struct {
}

// Dec implements BalanceRepository.
func (b balanceRepo) Dec(count int) error {
	panic("unimplemented")
}

// Get implements BalanceRepository.
func (b balanceRepo) Get(uID int) int {
	panic("unimplemented")
}

// GetWithdrawals implements BalanceRepository.
func (b balanceRepo) GetWithdrawals(uId int) []models.Withdrawal {
	panic("unimplemented")
}

// Inc implements BalanceRepository.
func (b balanceRepo) Inc(count int) error {
	panic("unimplemented")
}

func NewBalanceRepo(connection string) balanceRepository {
	return balanceRepo{}
}
