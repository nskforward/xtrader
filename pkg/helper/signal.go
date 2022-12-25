package helper

import (
	"os"
	"os/signal"
	"syscall"
)

func Signal() chan os.Signal {
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	return ch
}
