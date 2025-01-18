package models

type AccuralOrder struct {
	ID     int    `json:"order"`
	Status string `json:"status"`
	Count  string `json:"accural"`
}
