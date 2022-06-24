package noop

// Logger is a logger that writes nothing.
type Logger struct{}

// Trace logs a message at LevelTrace level.
func (*Logger) Trace(args ...interface{}) {}

// Tracef logs a message at LevelTrace level.
func (*Logger) Tracef(format string, args ...interface{}) {}

// Debug logs a message at LevelDebug level.
func (*Logger) Debug(args ...interface{}) {}

// Debugf logs a message at LevelDebug level.
func (*Logger) Debugf(format string, args ...interface{}) {}

// Info logs a message at LevelInfo level.
func (*Logger) Info(args ...interface{}) {}

// Infof logs a message at LevelInfo level.
func (*Logger) Infof(format string, args ...interface{}) {}

// Warn logs a message at LevelWarn level.
func (*Logger) Warn(args ...interface{}) {}

// Warnf logs a message at LevelWarn level.
func (*Logger) Warnf(format string, args ...interface{}) {}

// Error logs a message at LevelError level.
func (*Logger) Error(args ...interface{}) {}

// Errorf logs a message at LevelError level.
func (*Logger) Errorf(format string, args ...interface{}) {}
