package stream

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dihedron/seal/logging"
	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
)

const TimeFormat = "2006-01-02T15:04:05.999-0700"

// Logger is a logger that write sits messages to a stream.
type Logger struct {
	stream *os.File
}

// NewLogger returns an instance of a stream Logger.
func NewLogger(stream *os.File) *Logger {
	return &Logger{
		stream: stream,
	}
}

// Trace logs a message at LevelTrace level.
func (l *Logger) Trace(args ...interface{}) {
	if logging.GetLevel() <= logging.LevelTrace {
		frame := logging.GetCallerFrame(3)
		info := fmt.Sprintf("(%s:%d)", frame.File, frame.Line)
		if isatty.IsTerminal(l.stream.Fd()) {
			l.write(color.HiWhiteString("TRC"), append(args, info)...)
		} else {
			l.write("TRC", append(args, info)...)
		}
	}
}

// Tracef logs a message at LevelTrace level.
func (l *Logger) Tracef(msg string, args ...interface{}) {
	if logging.GetLevel() <= logging.LevelTrace {
		frame := logging.GetCallerFrame(3)
		info := fmt.Sprintf("(%s:%d)", frame.File, frame.Line)
		if isatty.IsTerminal(l.stream.Fd()) {
			l.writef(color.HiWhiteString("TRC"), msg+" "+info, args...)
		} else {
			l.writef("TRC", msg+" "+info, args...)
		}
	}
}

// Debug logs a message at LevelDebug level.
func (l *Logger) Debug(args ...interface{}) {
	if logging.GetLevel() <= logging.LevelDebug {
		frame := logging.GetCallerFrame(3)
		info := fmt.Sprintf("(%s:%d)", frame.File, frame.Line)
		if isatty.IsTerminal(l.stream.Fd()) {
			l.write(color.HiBlueString("DBG"), append(args, info)...)
		} else {
			l.write("DBG", append(args, info)...)
		}
	}
}

// Debugf logs a message at LevelDebug level.
func (l *Logger) Debugf(msg string, args ...interface{}) {
	if logging.GetLevel() <= logging.LevelDebug {
		frame := logging.GetCallerFrame(3)
		info := fmt.Sprintf("(%s:%d)", frame.File, frame.Line)
		if isatty.IsTerminal(l.stream.Fd()) {
			l.writef(color.HiBlueString("DBG"), msg+" "+info, args...)
		} else {
			l.writef("DBG", msg+" "+info, args...)
		}
	}
}

// Info logs a message at LevelInfo level.
func (l *Logger) Info(args ...interface{}) {
	if logging.GetLevel() <= logging.LevelInfo {
		frame := logging.GetCallerFrame(3)
		info := fmt.Sprintf("(%s:%d)", frame.File, frame.Line)
		if isatty.IsTerminal(l.stream.Fd()) {
			l.write(color.HiGreenString("INF"), append(args, info)...)
		} else {
			l.write("INF", append(args, info)...)
		}
	}
}

// Infof logs a message at LevelInfof level.
func (l *Logger) Infof(msg string, args ...interface{}) {
	if logging.GetLevel() <= logging.LevelInfo {
		frame := logging.GetCallerFrame(3)
		info := fmt.Sprintf("(%s:%d)", frame.File, frame.Line)
		if isatty.IsTerminal(l.stream.Fd()) {
			l.writef(color.HiGreenString("INF"), msg+" "+info, args...)
		} else {
			l.writef("INF", msg+" "+info, args...)
		}
	}
}

// Warn logs a message at LevelWarn level.
func (l *Logger) Warn(args ...interface{}) {
	if logging.GetLevel() <= logging.LevelWarn {
		frame := logging.GetCallerFrame(3)
		info := fmt.Sprintf("(%s:%d)", frame.File, frame.Line)
		if isatty.IsTerminal(l.stream.Fd()) {
			l.write(color.HiYellowString("WRN"), append(args, info)...)
		} else {
			l.write("WRN", append(args, info)...)
		}
	}
}

// Warnf logs a message at LevelWarn level.
func (l *Logger) Warnf(msg string, args ...interface{}) {
	if logging.GetLevel() <= logging.LevelWarn {
		frame := logging.GetCallerFrame(3)
		info := fmt.Sprintf("(%s:%d)", frame.File, frame.Line)
		if isatty.IsTerminal(l.stream.Fd()) {
			l.writef(color.HiYellowString("WRN"), msg+" "+info, args...)
		} else {
			l.writef("WRN", msg+" "+info, args...)
		}
	}
}

// Error logs a message at LevelError level.
func (l *Logger) Error(args ...interface{}) {
	if logging.GetLevel() <= logging.LevelError {
		frame := logging.GetCallerFrame(3)
		info := fmt.Sprintf("(%s:%d)", frame.File, frame.Line)
		if isatty.IsTerminal(l.stream.Fd()) {
			l.write(color.HiRedString("ERR"), append(args, info)...)
		} else {
			l.write("ERR", append(args, info)...)
		}
	}
}

// Errorf logs a message at LevelError level.
func (l *Logger) Errorf(msg string, args ...interface{}) {
	if logging.GetLevel() <= logging.LevelError {
		frame := logging.GetCallerFrame(3)
		info := fmt.Sprintf("(%s:%d)", frame.File, frame.Line)
		if isatty.IsTerminal(l.stream.Fd()) {
			l.writef(color.HiRedString("ERR"), msg+" "+info, args...)
		} else {
			l.writef("ERR", msg+" "+info, args...)
		}
	}
}

func (l *Logger) write(level string, args ...interface{}) {
	var buffer bytes.Buffer
	for argNum, arg := range args {
		if argNum > 0 {
			buffer.WriteString(" ")
		}
		buffer.WriteString(fmt.Sprintf("%v", arg))
	}
	fmt.Fprintf(l.stream, "%s [%s] %s\n", time.Now().Format(TimeFormat), level, buffer.String())
}

func (l *Logger) writef(level string, msg string, args ...interface{}) {
	message := fmt.Sprintf(strings.TrimSpace(msg), args...)
	fmt.Fprintf(l.stream, "%s [%s] %s\n", time.Now().Format(TimeFormat), level, message)
}
