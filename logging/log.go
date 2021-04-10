package logging

import (
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

// Trace logs a message at level Trace on the standard logger.
func Trace(message interface{}) {
	logrus.Debug(message)
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(message interface{}) {
	logrus.Debug(message)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func Info(message interface{}) {
	logrus.Info(message)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(message interface{}) {
	logrus.Warn(message)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(message interface{}) {
	logrus.Error(message)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(message interface{}) {
	logrus.Fatal(message)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}
