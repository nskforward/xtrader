package broker

import (
	"context"
	"errors"
	"fmt"
)

type Broker struct {
}

func New(ctx context.Context, cfg Config) (*Broker, error) {
	if cfg.Name == "" {
		return nil, fmt.Errorf("broker name must be defined in the config")
	}

	switch cfg.Name {
	case "exante":
		/*
			client, err := exante.NewClient(
				ctx, cfg.Exante.AccountID,
				cfg.Exante.Addr,
				cfg.Exante.ClientID,
				cfg.Exante.AppID,
				cfg.Exante.SharedKey,
			)
			if err != nil {
				return nil, fmt.Errorf("cannot connect to exante broker: %w", err)
			}
			return &Broker{ctx: ctx, client: client}, nil
		*/
		return nil, errors.New("not implemented")

	default:
		return nil, fmt.Errorf("unknown broker: %s", cfg.Name)
	}
}
