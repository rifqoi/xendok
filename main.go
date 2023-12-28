package main

import (
	"github.com/rifqoi/xendok-service/internal/logger"
)

func main() {
	l := logger.Get()
	defer l.Sync()

	l.Info("asd")
}
