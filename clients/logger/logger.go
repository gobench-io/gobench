package logger

import (
	"context"
	"log"

	"github.com/gobench-io/gobench/executor/driver"
)

const (
	// Topic default topic name
	Topic = "app-event-log"
	//DefaultName of event log
	DefaultName = "application"
	//DefaultSource of event log
	DefaultSource = "scenario"
	//DefaultLevel of event log
	DefaultLevel = "error"
)

// Logger represent to Logger
type Logger struct {
	ctx    context.Context
	source string
	level  string
	name   string
}

// NewLogger initial logger
func NewLogger(ctx context.Context, name, source, level string) *Logger {
	if name == "" {
		name = DefaultName
	}
	if source == "" {
		source = DefaultSource
	}
	if level == "" {
		level = DefaultLevel
	}

	return &Logger{
		ctx,
		source,
		name,
		level,
	}
}

// Log log an event
func (l *Logger) Log(message string) {
	driver.Notify(Topic, 1)
	log.Printf("")
}

func (l *Logger) Printf(format string, v ...interface{}) {
	driver.Notify(Topic, 1)
}
