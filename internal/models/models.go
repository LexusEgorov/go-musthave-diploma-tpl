package models

const (
	RegisteredStatus = "REGISTERED"
	InvalidStatus    = "INVALID"
	ProcessingStatus = "PROCESSING"
	ProcessedStatus  = "PROCESSED"
)

type AccuralOrder struct {
	ID     int    `json:"order"`
	Status string `json:"status"`
	Count  string `json:"accural"`
}

type User struct {
	Id       int
	Login    string
	Password string
}

type UserAuth struct {
	Jwt string
}

type Order struct {
	ID           int
	Number       int
	BonusesCount int
	Status       string
}

type Withdrawal struct {
	Order       int
	Sum         int
	ProcessedAt string
}
