package golang

import (
	"bytes"
	"fmt"
	golang "log"
	"os"
	"strings"

	"github.com/dihedron/seal/logging"
)

// Logger is te type wrapping the default Golang logger.
type Logger struct {
	logger *golang.Logger
}

// NewLogger returns a new Golang Logger.
func NewLogger(prefix string) *Logger {
	return &Logger{
		logger: golang.New(os.Stderr, prefix, golang.Ltime|golang.Ldate|golang.Lmicroseconds),
	}
}

func (l *Logger) Trace(args ...interface{}) {
	if logging.GetLevel() <= logging.LevelTrace {
		var buffer bytes.Buffer
		for argNum, arg := range args {
			if argNum > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(fmt.Sprintf("%v", arg))
		}
		message := fmt.Sprintf("[TRC] %s", buffer.String())
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Tracef(msg string, args ...interface{}) {
	if logging.GetLevel() <= logging.LevelTrace {
		message := fmt.Sprintf("[TRC] "+msg, args...)
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Debug(args ...interface{}) {
	if logging.GetLevel() <= logging.LevelDebug {
		var buffer bytes.Buffer
		for argNum, arg := range args {
			if argNum > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(fmt.Sprintf("%v", arg))
		}
		message := fmt.Sprintf("[DBG] %s", buffer.String())
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	if logging.GetLevel() <= logging.LevelDebug {
		message := fmt.Sprintf("[DBG] "+msg, args...)
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Info(args ...interface{}) {
	if logging.GetLevel() <= logging.LevelInfo {
		var buffer bytes.Buffer
		for argNum, arg := range args {
			if argNum > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(fmt.Sprintf("%v", arg))
		}
		message := fmt.Sprintf("[INF] %s", buffer.String())
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	if logging.GetLevel() <= logging.LevelInfo {
		message := fmt.Sprintf("[INF] "+msg, args...)
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Warn(args ...interface{}) {
	if logging.GetLevel() <= logging.LevelWarn {
		var buffer bytes.Buffer
		for argNum, arg := range args {
			if argNum > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(fmt.Sprintf("%v", arg))
		}
		message := fmt.Sprintf("[WRN] %s", buffer.String())
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	if logging.GetLevel() <= logging.LevelWarn {
		message := fmt.Sprintf("[WRN] "+msg, args...)
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Error(args ...interface{}) {
	if logging.GetLevel() <= logging.LevelError {
		var buffer bytes.Buffer
		for argNum, arg := range args {
			if argNum > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(fmt.Sprintf("%v", arg))
		}
		message := fmt.Sprintf("[ERR] %s", buffer.String())
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	if logging.GetLevel() <= logging.LevelError {
		message := fmt.Sprintf("[ERR] "+msg, args...)
		message = strings.TrimRight(message, "\n\r")
		golang.Print(message)
	}
}
