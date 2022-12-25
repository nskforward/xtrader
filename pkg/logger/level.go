package logger

import "fmt"

type Level uint8

const (
	debug Level = iota
	info
	warn
	err
	fatal
)

func (level Level) prefix() string {
	switch level {
	case debug:
		return "[debug]"
	case info:
		return "[info]"
	case warn:
		return "[warn]"
	case err:
		return "[error]"
	case fatal:
		return "[fatal]"
	default:
		panic(fmt.Errorf("unknown logger level: %d", level))
	}
}
