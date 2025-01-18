package order

type order struct {
	id uint
}

func NewOrder(id uint) *order {
	return &order{
		id: id,
	}
}
