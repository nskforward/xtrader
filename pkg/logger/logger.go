package logger

import (
	"fmt"
	"io"
	"syscall"
)

type Logger struct {
	out   io.Writer
	err   io.Writer
	debug bool
	time  bool
}

func New(debug, time bool, out, err io.Writer) *Logger {
	return &Logger{
		debug: debug,
		time:  time,
		out:   out,
		err:   err,
	}
}

func (l *Logger) stream(level Level) io.Writer {
	switch level {
	case debug, info:
		return l.out
	case warn, err, fatal:
		return l.err
	default:
		panic(fmt.Errorf("unknown logger level: %d", level))
	}
}

func (l *Logger) Debug(msg ...any) {
	if !l.debug {
		return
	}
	err := l.write(debug, msg...)
	if err != nil {
		panic(fmt.Errorf("cannot write log debug message: %w", err))
	}
}

func (l *Logger) Info(msg ...any) {
	err := l.write(info, msg...)
	if err != nil {
		panic(fmt.Errorf("cannot write log info message: %w", err))
	}
}

func (l *Logger) Warning(msg ...any) {
	err := l.write(warn, msg...)
	if err != nil {
		panic(fmt.Errorf("cannot write log warn message: %w", err))
	}
}

func (l *Logger) Error(msg ...any) {
	err := l.write(err, msg...)
	if err != nil {
		panic(fmt.Errorf("cannot write log err message: %w", err))
	}
}

func (l *Logger) Fatal(msg ...any) {
	err := l.write(fatal, msg...)
	if err != nil {
		panic(fmt.Errorf("cannot write log fatal message: %w", err))
	}
	err = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	if err != nil {
		panic(fmt.Errorf("cannot send interrupt signal: %w", err))
	}
}
