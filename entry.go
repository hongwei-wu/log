package log

import (
	"github.com/hongwei-wu/log/field"
	"github.com/hongwei-wu/log/util"
	"github.com/sirupsen/logrus"
)

const (
	Skip = 4
)

type Entry struct {
	entry *logrus.Entry
}

func NewEntry(logger *logrus.Logger) *Entry {
	return &Entry{entry: logrus.NewEntry(logger)}
}

func (e *Entry) withFileField() *Entry {
	if _, ok := e.entry.Data[field.File]; ok {
		return e
	}

	var skip int
	if v, ok := e.entry.Data[field.Skip]; ok {
		skip, _ = v.(int)
	}
	return e.WithField(field.File, util.FileAndLine(Skip+skip))
}

func (e *Entry) WithField(key string, value interface{}) *Entry {
	return &Entry{entry: e.entry.WithField(key, value)}
}

func (e *Entry) WithFields(fields logrus.Fields) *Entry {
	return &Entry{entry: e.entry.WithFields(fields)}
}

func (e *Entry) Debugf(format string, args ...interface{}) {
	e.printf(logrus.DebugLevel, format, args...)
}

func (e *Entry) Infof(format string, args ...interface{}) {
	e.printf(logrus.InfoLevel, format, args...)
}

func (e *Entry) Warnf(format string, args ...interface{}) {
	e.printf(logrus.WarnLevel, format, args...)
}

func (e *Entry) Warningf(format string, args ...interface{}) {
	e.printf(logrus.WarnLevel, format, args...)
}

func (e *Entry) Errorf(format string, args ...interface{}) {
	e.printf(logrus.ErrorLevel, format, args...)
}

func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.printf(logrus.FatalLevel, format, args...)
}

func (e *Entry) printf(level logrus.Level, format string, args ...interface{}) {
	e.withFileField().entry.Logf(level, format, args...)
}

func (e *Entry) Debug(args ...interface{}) {
	e.print(logrus.DebugLevel, args...)
}

func (e *Entry) Info(args ...interface{}) {
	e.print(logrus.InfoLevel, args...)
}

func (e *Entry) Warn(args ...interface{}) {
	e.print(logrus.WarnLevel, args...)
}

func (e *Entry) Warning(args ...interface{}) {
	e.print(logrus.WarnLevel, args...)
}

func (e *Entry) Error(args ...interface{}) {
	e.print(logrus.ErrorLevel, args...)
}

func (e *Entry) Fatal(args ...interface{}) {
	e.print(logrus.FatalLevel, args...)
}

func (e *Entry) print(level logrus.Level, args ...interface{}) {
	e.withFileField().entry.Log(level, args...)
}

func (e *Entry) Debugln(args ...interface{}) {
	e.println(logrus.DebugLevel, args...)
}

func (e *Entry) Infoln(args ...interface{}) {
	e.println(logrus.InfoLevel, args...)
}

func (entry *Entry) Warnln(args ...interface{}) {
	entry.println(logrus.WarnLevel, args...)
}

func (e *Entry) Warningln(args ...interface{}) {
	e.println(logrus.WarnLevel, args...)
}

func (entry *Entry) Errorln(args ...interface{}) {
	entry.println(logrus.ErrorLevel, args...)
}

func (e *Entry) Fatalln(args ...interface{}) {
	e.println(logrus.FatalLevel, args...)
}

func (e *Entry) println(level logrus.Level, args ...interface{}) {
	e.withFileField().entry.Logln(level, args...)
}
