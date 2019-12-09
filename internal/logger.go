package internal

import (
	"encoding/binary"
	"github.com/hongwei-wu/log/internal/appender"
	"github.com/hongwei-wu/log/internal/formatter"
	"github.com/hongwei-wu/log/internal/opts"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	once      sync.Once
	defLogger *Logger
)

func GetDefaultLogger() *Logger {
	once.Do(func() {
		defLogger = NewLogger()
		var opts opts.Opts
		opts = append(opts, SetLevelOpt(defLogger, DEBUG))
		opts = append(opts, SetFormatterOpt(defLogger, formatter.CreateFormatter(formatter.Raw)))
		opts = append(opts, SetAppenderOpt(defLogger, appender.CreateAppender(appender.Console)))
		defLogger.Apply(opts)
	})
	return defLogger
}

type Logger struct {
	*logrus.Logger
	level     string
	appenders []appender.Appender
}

func NewLogger() *Logger {
	l := &Logger{
		Logger: logrus.New(),
	}
	l.Logger.Out = l
	return l
}

func (l *Logger) Write(p []byte) (n int, err error) {
	level := binary.LittleEndian.Uint32(p)
	for _, a := range l.appenders {
		if uint32(a.GetLevel()) >= level {
			a.Write(p[4:])
		}
	}
	return len(p), nil
}

func (l *Logger) Apply(opts opts.Opts) error {
	return opts.Apply(l)
}

func (l *Logger) SetLevel(level string) {
	l.Logger.SetLevel(ParseLevel(level))
}

func SetLevelOpt(ptr interface{}, level string) opts.Opt {
	return func() error {
		l := ptr.(*Logger)
		l.level = level
		l.SetLevel(level)
		return nil
	}
}

func SetFormatterOpt(ptr interface{}, formatter formatter.Formatter) opts.Opt {
	return func() error {
		l := ptr.(*Logger)
		l.Logger.Formatter = formatter
		return nil
	}
}

func SetAppenderOpt(ptr interface{}, appender appender.Appender) opts.Opt {
	return func() error {
		l := ptr.(*Logger)
		l.appenders = append(l.appenders, appender)
		return nil
	}
}
