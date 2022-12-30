package main

import (
	"context"
	"github.com/nskforward/xtrader/internal/broker"
	"github.com/nskforward/xtrader/internal/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Get()

	_, err := broker.New(ctx, cfg.Broker)
	if err != nil {
		panic(err)
	}
}
