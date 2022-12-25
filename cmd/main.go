package main

import (
	"context"
	"github.com/nskforward/xtrader/internal/broker/exante"
	"github.com/nskforward/xtrader/internal/config"
	"github.com/nskforward/xtrader/internal/httpclient"
	"github.com/nskforward/xtrader/pkg/logger"
	"os"
	"time"
)

func main() {
	l := logger.New(true, true, os.Stdout, os.Stderr)
	httpclient.InitDefault()

	l.Info("service started")
	defer l.Info("service stopped")

	cfg := config.Get()
	l.Info("broker:", cfg.Broker.Name)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ch, err := exante.WatchQuotes(ctx, cfg.Broker.Exante, "BTC.USD")
	if err != nil {
		l.Fatal(err)
		return
	}

LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case q, ok := <-ch:
			if !ok {
				l.Info("channel closed")
				break LOOP
			}
			price, size, err := q.Bid_()
			if err != nil {
				l.Error(err)
			} else {
				l.Info(q.Symbol_(), "bid:", price, size)
			}
		}
	}
}
