package balance

type balance struct {
	uId     uint
	balance int
}

func NewBalance(uId uint) *balance {
	return &balance{
		uId:     uId,
		balance: 0,
	}
}
