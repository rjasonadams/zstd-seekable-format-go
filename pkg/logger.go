package seekable

// Logger is a minimal logging interface for the library.
// Implementations can wrap any logging library (zap, logrus, slog, etc.).
type Logger interface {
	// Debug logs a debug-level message with structured fields.
	Debug(msg string, fields ...Field)
}

// Field represents a structured logging field with a key-value pair.
type Field struct {
	Key   string
	Value interface{}
}

// String creates a Field with a string value.
func String(key, value string) Field {
	return Field{Key: key, Value: value}
}

// Uint64 creates a Field with a uint64 value.
func Uint64(key string, value uint64) Field {
	return Field{Key: key, Value: value}
}

// Int creates a Field with an int value.
func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

// Any creates a Field with any value type.
func Any(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

// NoOpLogger is a logger implementation that discards all log messages.
// This is the default logger used when none is provided.
type NoOpLogger struct{}

// Debug implements Logger.Debug by doing nothing.
func (NoOpLogger) Debug(string, ...Field) {}
