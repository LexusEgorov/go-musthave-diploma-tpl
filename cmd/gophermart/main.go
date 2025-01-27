package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}
