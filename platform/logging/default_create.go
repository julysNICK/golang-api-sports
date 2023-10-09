package logging

import (
	"log"
	"os"
	"platform/config"
	"strings"
)

func NewDefaultLogger(cfg config.Configuration) Logger {

	var level LogLevel = Debug

	if configLevelString, found := cfg.GetString("logging:level"); found {
		level = LogLevelFromString(configLevelString)
	}
	flags := log.Lmsgprefix | log.Ltime

	return &DefaultLogger{
		minLevel: level,
		loggers: map[LogLevel]*log.Logger{
			Trace:       log.New(log.Writer(), "TRACE: ", flags),
			Debug:       log.New(log.Writer(), "DEBUG: ", flags),
			Information: log.New(log.Writer(), "INFO: ", flags),
			Warning:     log.New(os.Stdout, "WARN: ", flags),
			Fatal:       log.New(os.Stdout, "FATAL: ", flags),
		},
		triggersPanics: true,
	}
}

func LogLevelFromString(val string) (level LogLevel) {
	switch strings.ToLower(val) {
	case "debug":
		level = Debug
	case "information":
		level = Information
	case "warning":
		level = Warning
	case "fatal":
		level = Fatal
	case "none":
		level = None
	default:
		level = Debug
	}
	return
}