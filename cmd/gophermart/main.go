package main

import (
	"fmt"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/db"
)

const (
	host     = "localhost"
	port     = 53322
	user     = "root"
	password = "root"
	dbname   = "root"
)

func main() {
	// log := logrus.New()

	// log.SetFormatter(&logrus.TextFormatter{
	// 	ForceColors:   true,
	// 	FullTimestamp: true,
	// })
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db.NewDB(psqlInfo, false)
}
