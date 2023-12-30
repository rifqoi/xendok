package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rifqoi/xendok-service/internal/config"
	"github.com/rifqoi/xendok-service/internal/logger"
)

func main() {

	args := config.ProcessArgs()
	err := config.Init(args)
	if err != nil {
		log.Panicf("failed to init config: %v", err)
	}

	l := logger.Get()
	defer l.Sync()

	ctx := logger.WithCtx(context.Background(), l)

	cfg := config.Get()

	fmt.Println(cfg.Database.Password)
	logSomething(ctx)
	l.Info("asd")
	l.Debug("asd")
}

func logSomething(ctx context.Context) {
	l := logger.FromCtx(ctx)
	l.Info("asdasdad")
}
