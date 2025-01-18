package client

import "github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"

type client struct {
}

func (c client) SendRequest(host string) *models.AccuralOrder {
	return nil
}

func NewClient() *client {
	return &client{}
}
