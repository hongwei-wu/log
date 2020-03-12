package log

import "github.com/sirupsen/logrus"

const (
	DEBUG = "debug"
	INFO  = "info"
	WARN  = "warn"
	ERROR = "error"
	FATAL = "fatal"
)

func ParseLevel(level string) logrus.Level {
	if level == DEBUG {
		return logrus.DebugLevel
	} else if level == INFO {
		return logrus.InfoLevel
	} else if level == WARN {
		return logrus.WarnLevel
	} else if level == ERROR {
		return logrus.ErrorLevel
	} else if level == FATAL {
		return logrus.FatalLevel
	} else {
		return logrus.InfoLevel
	}
}
