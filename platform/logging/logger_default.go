package logging

import (
	"fmt"
	"log"
)

type DefaultLogger struct {
	minLevel LogLevel

	loggers map[LogLevel]*log.Logger

	triggersPanics bool
}

func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

func (l *DefaultLogger) Write(level LogLevel, message string) {
	if l.minLevel <= level {
		l.loggers[level].Output(2, message)
	}
}

func (l *DefaultLogger) Trace(msg string) {
	l.Write(Trace, msg)
}

func (l *DefaultLogger) Tracef(template string, vals ...interface{}) {
	l.Write(Debug, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Debug(msg string) {
	l.Write(Debug, msg)
}

func (l *DefaultLogger) Debugf(template string, vals ...interface{}) {
	l.Write(Debug, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Info(msg string) {
	l.Write(Information, msg)
}

func (l *DefaultLogger) Infof(template string, vals ...interface{}) {
	l.Write(Information, fmt.Sprintf(template, vals...))
}
func (l *DefaultLogger) Warn(msg string) {
	l.Write(Warning, msg)
}

func (l *DefaultLogger) Warnf(template string, vals ...interface{}) {
	l.Write(Warning, fmt.Sprintf(template, vals...))
}

func (l *DefaultLogger) Panic(msg string) {
	l.Write(Fatal, msg)
	if l.triggersPanics {
		panic(msg)
	}
}

func (l *DefaultLogger) Panicf(template string, vals ...interface{}) {
	formatted := fmt.Sprintf(template, vals...)
	l.Write(Fatal, formatted)

	if l.triggersPanics {
		panic(formatted)
	}
}

func (l *DefaultLogger) Warning(msg string) {
	l.Write(Warning, msg)
}

func (l *DefaultLogger) Warningf(template string, vals ...interface{}) {
	l.Write(Warning, fmt.Sprintf(template, vals...))
}
