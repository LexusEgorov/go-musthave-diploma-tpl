package main

import (
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/config"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/server"
	"github.com/go-chi/chi"
)

func main() {
	conf := config.NewConfig()
	serv := server.NewServer(conf, chi.NewRouter())

	serv.Serve()
}
