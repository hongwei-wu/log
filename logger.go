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

func (impl *LoggerImpl) Write(p []byte) (n int, err error) {
	for _, a := range impl.appenders {
		a.Write(p)
	}
	return len(p), nil
}

func (impl *LoggerImpl) SetLevel(level string) {
	impl.Logger.SetLevel(ParseLevel(level))
}

func (impl *LoggerImpl) SetFormatter(formatter formatter.Formatter) {
	impl.Logger.Formatter = formatter
}

func (impl *LoggerImpl) AddAppender(appender appender.Appender) {
	impl.appenders[appender.Name()] = appender
}

func (impl *LoggerImpl) ResetAppender() {
	impl.appenders = make(map[string]appender.Appender)
}
