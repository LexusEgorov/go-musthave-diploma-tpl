package models

const (
	RegisteredStatus = "REGISTERED"
	InvalidStatus    = "INVALID"
	ProcessingStatus = "PROCESSING"
	ProcessedStatus  = "PROCESSED"
)

type AccuralOrder struct {
	ID     string  `json:"order"`
	Status string  `json:"status"`
	Count  float64 `json:"accural"`
}

type User struct {
	Login    string
	Password string
}

type UserAuth struct {
	Jwt string
}

type Order struct {
	Number       int     `json:"number"`
	BonusesCount float64 `json:"accural,omitempty"`
	Status       string  `json:"status"`
	CreatedAt    string  `json:"uploaded_at"`
}

type Withdrawal struct {
	Order       int     `json:"order"`
	Sum         float64 `json:"sum"`
	ProcessedAt string  `json:"processed_at"`
}

type UserBalance struct {
	Balance   float64 `json:"current"`
	BalanceWD float64 `json:"withdrawn"`
}

type WdBalance struct {
	Order string  `json:"order"`
	Sum   float64 `json:"sum"`
}

type UserUpdate struct {
	ID    int
	Count float64
}
