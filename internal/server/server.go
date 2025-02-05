package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/balance"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/client"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/config"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/db"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/order"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/server/handlers"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/server/middleware"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/user"
)

type server struct {
	client     client.Client
	host       string
	router     *chi.Mux
	httpServer *http.Server
}

func (s *server) Serve() error {
	s.httpServer = &http.Server{
		Addr:    s.host,
		Handler: s.router,
	}

	s.client.Run()
	logrus.Info("Server is running on: ", s.host)
	return s.httpServer.ListenAndServe()
}

func (s *server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("Error while stopping: %s\n", err)
	}
}

func NewServer(config config.Server, router *chi.Mux) *server {
	dbProvider := db.NewDB(config.DB, false)

	userRepo := user.NewUserRepo(*dbProvider)
	user := user.NewUser(userRepo)

	balanceRepo := balance.NewBalanceRepo(*dbProvider)
	balance := balance.NewBalance(balanceRepo)

	orderRepo := order.NewOrderRepo(*dbProvider)
	order := order.NewOrder(orderRepo)

	handlers := handlers.NewHandlers(user, balance, order)

	serverClient := client.NewClient(order, config.Host)

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

	s := &server{
		client: *serverClient,
		host:   config.Host,
		router: router,
	}

	return s
}
