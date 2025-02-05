package order

import (
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
	servererrors "github.com/LexusEgorov/go-musthave-diploma-tpl/internal/serverErrors"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/utils"
	"github.com/sirupsen/logrus"
)

type orderRepository interface {
	Add(uID int, o string, count float64) error
	Get(uID int) ([]models.Order, error)
	GetOrder(o string) (int, error)
	GetWithdrawals(uID int) ([]models.Withdrawal, error)
	Update(o models.AccuralOrder) (*models.UserUpdate, error)
	GetQueue() []string
}

type order struct {
	repo orderRepository
}

// GetQueue implements client.orderManager.
func (o order) GetQueue() []string {
	return o.repo.GetQueue()
}

func (o order) Update(order models.AccuralOrder) (*models.UserUpdate, error) {
	return o.repo.Update(order)
}

// Add implements handlers.OrderManager.
func (o order) Add(uID int, number string, wdSum *float64) error {
	if !utils.LunaCheck(number) {
		return servererrors.WrongNumberError{Number: number}
	}

	user, err := o.repo.GetOrder(number)

	if err != nil {
		logrus.Error(err)
		return err
	}

	if user == 0 {
		var sum float64 = 0

		if wdSum != nil {
			sum = *wdSum * -1
		}

		return o.repo.Add(uID, number, sum)
	}

	if user == uID {
		return servererrors.OkayError{}
	}

	return servererrors.ConflictError{Used: number}
}

// Get implements handlers.OrderManager.
func (o order) Get(uID int) []models.Order {
	order, err := o.repo.Get(uID)

	if err != nil {
		logrus.Error(err)
		return nil
	}

	return order
}

// GetWithdrawals implements handlers.OrderManager.
func (o order) GetWithdrawals(uID int) []models.Withdrawal {
	withdrawals, err := o.repo.GetWithdrawals(uID)

	if err != nil {
		logrus.Error(err)
	}

	return withdrawals
}

func NewOrder(repo orderRepository) order {
	return order{
		repo: repo,
	}
}
