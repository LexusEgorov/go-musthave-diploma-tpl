package models

type AccuralOrder struct {
	ID     int    `json:"order"`
	Status string `json:"status"`
	Count  string `json:"accural"`
}

type User struct {
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
