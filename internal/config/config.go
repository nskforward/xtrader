package config

import (
	"encoding/json"
	"fmt"
	"github.com/nskforward/xtrader/internal/broker"
	"github.com/nskforward/xtrader/pkg/helper"
	"io/ioutil"
)

type Config struct {
	Broker broker.Config `json:"broker"`
}

func newConfig() *Config {
	data, err := ioutil.ReadFile(helper.ResolvePath("config.json"))
	if err != nil {
		panic(fmt.Errorf("cannot read config.json file: %w", err))
	}
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		panic(fmt.Errorf("cannot unmarshal config.json: %w", err))
	}
	return &cfg
}
