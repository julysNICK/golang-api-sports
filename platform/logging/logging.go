package logging

type LogLevel int

const (
	Trace LogLevel = iota

	Debug
	Information
	Warning
	Fatal
	None
)

type Logger interface {
	Trace(message string)

	Tracef(format string, args ...interface{})

	Debug(message string)
	Debugf(format string, args ...interface{})

	Info(message string)
	Infof(format string, args ...interface{})

	Warning(message string)
	Warningf(format string, args ...interface{})

	Panic(message string)
	Panicf(format string, args ...interface{})
}
