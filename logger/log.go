package logger

import (
	"go.uber.org/zap"
)

// Logger interface of the gobench server
type Logger interface {
	// Logs an information statement
	Infow(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	Sync() error
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

// NewNopLogger returns a no-op Logger. It never writes out logs or internal errors,
// and it never runs user-defined hooks.
func NewNopLogger() *Log {
	nopLogger := zap.NewNop().Sugar()
	return &Log{
		SugaredLogger: *nopLogger,
	}
}

// NewApplicationLogger setup log when running an application
func NewApplicationLogger(f string) (*Log, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = append(cfg.OutputPaths, f)
	log, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return &Log{
		SugaredLogger: *log.Sugar(),
	}, nil
}
