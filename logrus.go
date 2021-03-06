package logrus

import (
	"fmt"

	"github.com/marcaudefroy/loggers"
	"github.com/sirupsen/logrus"
)

// Logger is an Contextual logger wrapper over Logrus's logger.
type Logger struct {
	*logrus.Logger
}

// NewLogger returns a Contextual Logger for Logrus's logger.
// Note that any initialization must be done on the input logrus.
func NewLogger(logrus *logrus.Logger) loggers.Contextual {
	var l Logger
	l.Logger = logrus
	l.Info("Now using Logrus logger package (via loggers/mappers/logrus).")
	return &l
}

// NewDefaultLogger returns a Contextual Logger for Logrus's logger.
// The logger will contain whatever defaults Logrus uses.
func NewDefaultLogger() loggers.Contextual {
	var l Logger
	l.Logger = logrus.New()
	l.Info("Now using Logrus logger package (via loggers/mappers/logrus).")
	return &l
}

// WithField returns an advanced logger with a pre-set field.
func (l *Logger) WithField(key string, value interface{}) loggers.Contextual {
	entry := &LoggerEntry{logger: l, Entry: l.Logger.WithField(key, value)}
	return entry
}

// WithFields returns an advanced logger with pre-set fields.
func (l *Logger) WithFields(fields ...interface{}) loggers.Contextual {
	entry := &LoggerEntry{logger: l, Entry: l.Logger.WithFields(sliceToMap(fields...))}
	return entry
}

type LoggerEntry struct {
	logger *Logger
	*logrus.Entry
}

// WithField returns an advanced logger with a pre-set field.
func (l *LoggerEntry) WithField(key string, value interface{}) loggers.Contextual {
	l.Entry = l.Entry.WithField(key, value)
	return l
}

// WithFields returns an advanced logger with a pre-set field.
func (l *LoggerEntry) WithFields(fields ...interface{}) loggers.Contextual {

	l.Entry = l.Entry.WithFields(sliceToMap(fields...))
	return l
}

func sliceToMap(fields ...interface{}) map[string]interface{} {
	f := make(map[string]interface{}, len(fields)/2)
	var key, value interface{}
	for i := 0; i+1 < len(fields); i = i + 2 {
		key = fields[i]
		value = fields[i+1]
		if s, ok := key.(string); ok {
			f[s] = value
		} else if s, ok := key.(fmt.Stringer); ok {
			f[s.String()] = value
		}
	}
	return f
}
