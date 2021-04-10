package at

import (
	"abelce/at/logging"
)

// Trace logs a message at level Trace on the standard logger.
func Trace(message interface{}) {
	logging.Debug(message)
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	logging.Debugf(format, args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(message interface{}) {
	logging.Debug(message)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	logging.Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func Info(message interface{}) {
	logging.Info(message)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	logging.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(message interface{}) {
	logging.Warn(message)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	logging.Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(message interface{}) {
	logging.Error(message)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	logging.Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(message interface{}) {
	logging.Fatal(message)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	logging.Fatalf(format, args...)
}
