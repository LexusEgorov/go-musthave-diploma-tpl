package config

import (
	"flag"
	"os"
)

type Server struct {
	Host            string
	DB              string
	AccuralHost     string
	IsCreatedTables bool
}

func NewConfig() Server {
	var host, db, accuralHost string

	flag.StringVar(&host, "a", "localhost:3000", "host of server")
	flag.StringVar(&db, "d", "", "connection string to db")
	flag.StringVar(&accuralHost, "r", "", "host of accural service")

	flag.Parse()

	if env := os.Getenv("RUN_ADDRESS"); env != "" {
		host = env
	}

	if env := os.Getenv("DATABASE_URI"); env != "" {
		db = env
	}

	if env := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); env != "" {
		accuralHost = env
	}

	return Server{
		Host:        host,
		DB:          db,
		AccuralHost: accuralHost,
	}
}
