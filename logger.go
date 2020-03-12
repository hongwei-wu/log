package log

import (
	"github.com/hongwei-wu/log/appender"
	"github.com/hongwei-wu/log/formatter"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	once       sync.Once
	rootLogger *LoggerImpl
)

func RootLogger() *LoggerImpl {
	once.Do(func() {
		rootLogger = NewLoggerImpl()
		rootLogger.SetFormatter(formatter.NewRawFormatter())
		rootLogger.AddAppender(appender.NewConsoleAppender())
	})
	return rootLogger
}

type LoggerImpl struct {
	*logrus.Logger
	appenders map[string]appender.Appender
}

func NewLoggerImpl() *LoggerImpl {
	l := &LoggerImpl{Logger: logrus.New(), appenders: make(map[string]appender.Appender)}
	l.Logger.Out = l
	return l
}

func (l *LoggerImpl) Write(p []byte) (n int, err error) {
	for _, a := range l.appenders {
		a.Write(p)
	}
	return len(p), nil
}

func (l *LoggerImpl) SetLevel(level string) {
	l.Logger.SetLevel(ParseLevel(level))
}

func (l *LoggerImpl) SetFormatter(formatter formatter.Formatter) {
	l.Logger.Formatter = formatter
}

func (l *LoggerImpl) AddAppender(appender appender.Appender) {
	l.appenders[appender.Name()] = appender
}
