package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/config"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/server/handlers"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/server/middleware"
)

type server struct {
	host       string
	router     *chi.Mux
	httpServer *http.Server
}

func (s server) Serve() error {
	s.httpServer = &http.Server{
		Addr:    s.host,
		Handler: s.router,
	}

	return s.httpServer.ListenAndServe()
}

func (s server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("Error while stopping: %s\n", err)
	}
}

func NewServer(config config.Server, router *chi.Mux) *server {
	handlers := handlers.NewHandlers()

	s := &server{
		host:   config.Host,
		router: router,
	}

	router.Use(middleware.WithDecoding)
	router.Use(middleware.WithEncoding)

	router.Route("/api/user", func(r chi.Router) {
		r.Post("/register", http.HandlerFunc(handlers.Registration))
		r.Post("/login", http.HandlerFunc(handlers.Auth))

		r.Group(func(r chi.Router) {
			r.Use(middleware.WithAuth)

			r.Post("/orders", http.HandlerFunc(handlers.AddOrders))
			r.Get("/orders", http.HandlerFunc(handlers.GetOrders))

			r.Route("/balance", func(r chi.Router) {
				r.Get("/", http.HandlerFunc(handlers.GetBalance))
				r.Post("/withdraw", http.HandlerFunc(handlers.WithdrawBalance))
			})

			r.Get("/withdrawals", http.HandlerFunc(handlers.GetWithdrawals))
		})
	})

	return s
}
