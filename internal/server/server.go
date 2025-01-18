package server

import (
	"github.com/go-chi/chi"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/config"
)

type server struct{}

func (s server) Serve() {}

func (s server) Stop() {}

func NewServer(config config.Server, router *chi.Mux) *server {
	return &server{}
}
