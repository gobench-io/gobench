package logger

import "go.uber.org/zap"

// Logger interface of the gobench server
type Logger interface {
	// Logs an information statement
	Infow(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
}

// Log is the wrap above zap sugar logger
type Log struct {
	zap.SugaredLogger
}

// NewStdLogger returns a zap sugar logger
func NewStdLogger() *Log {
	zapLogger, _ := zap.NewProduction()
	l := &Log{
		SugaredLogger: *zapLogger.Sugar(),
	}
	return l
}
