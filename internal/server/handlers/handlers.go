package handlers

type handlers struct{}

func (h handlers) Registration() {}
func (h handlers) Auth()         {}

func (h handlers) GetOrders() {}
func (h handlers) AddOrders() {}

func (h handlers) GetBalance()      {}
func (h handlers) WithdrawBalance() {}

func (h handlers) GetWithdrawals() {}

func NewHandlers() handlers {
	return handlers{}
}
