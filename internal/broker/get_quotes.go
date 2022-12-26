package broker

import (
	"context"
	"fmt"
	"github.com/nskforward/xtrader/internal/broker/exante"
	"github.com/nskforward/xtrader/internal/config"
	"github.com/nskforward/xtrader/internal/quote"
)

func GetQuotes(ctx context.Context, symbol string) (quotes chan quote.Quote, err error) {
	cfg := config.Get()
	switch cfg.Broker.Name {
	case "exante":
		return exante.GetQuotes(ctx, cfg.Broker.Exante, symbol)
	default:
		err = fmt.Errorf("unknown broker name: %s", cfg.Broker.Name)
		return
	}
}
