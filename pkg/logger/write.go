package logger

import (
	"fmt"
	"time"
)

func (l *Logger) write(level Level, messages ...any) error {

	w := l.stream(level)

	// write time
	if l.time {
		_, err := fmt.Fprint(w, time.Now().Format("2006/01/02 15:04:05.000 "))
		if err != nil {
			return err
		}
	}

	// write prefix
	_, err := fmt.Fprint(w, level.prefix())
	if err != nil {
		return err
	}

	// write message
	for _, msg := range messages {
		_, err := fmt.Fprint(w, " ")
		if err != nil {
			return err
		}
		_, err = fmt.Fprint(w, msg)
		if err != nil {
			return err
		}
	}

	// write new line sign
	_, err = fmt.Fprint(w, "\n")
	if err != nil {
		return err
	}

	return nil
}
