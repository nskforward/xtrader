package main

import (
	"github.com/nskforward/xtrader/pkg/logger"
	"os"
)

func main() {
	l := logger.New(true, true, os.Stdout, os.Stderr)
	l.Info("service started")
	defer l.Info("service stopped")

}
