package log

import (
	"github.com/sirupsen/logrus"
)

func WithField(key string, value interface{}) *Entry {
	return NewEntry(RootLogger().Logger).WithField(key, value)
}

func WithFields(fields logrus.Fields) *Entry {
	return NewEntry(RootLogger().Logger).WithFields(fields)
}

func SetLevel(level string)                       { RootLogger().SetLevel(level) }
func Debug(args ...interface{})                   { print(logrus.DebugLevel, args...) }
func Debugln(args ...interface{})                 { println(logrus.DebugLevel, args...) }
func Debugf(format string, args ...interface{})   { printf(logrus.DebugLevel, format, args...) }
func Info(args ...interface{})                    { print(logrus.InfoLevel, args...) }
func Infoln(args ...interface{})                  { println(logrus.InfoLevel, args...) }
func Infof(format string, args ...interface{})    { printf(logrus.InfoLevel, format, args...) }
func Warn(args ...interface{})                    { print(logrus.WarnLevel, args...) }
func Warnln(args ...interface{})                  { println(logrus.WarnLevel, args...) }
func Warnf(format string, args ...interface{})    { printf(logrus.WarnLevel, format, args...) }
func Warning(args ...interface{})                 { print(logrus.WarnLevel, args...) }
func Warningln(args ...interface{})               { println(logrus.WarnLevel, args...) }
func Warningf(format string, args ...interface{}) { printf(logrus.WarnLevel, format, args...) }
func Error(args ...interface{})                   { print(logrus.ErrorLevel, args...) }
func Errorln(args ...interface{})                 { println(logrus.ErrorLevel, args...) }
func Errorf(format string, args ...interface{})   { printf(logrus.ErrorLevel, format, args...) }
func Fatal(args ...interface{})                   { print(logrus.FatalLevel, args...) }
func Fatalln(args ...interface{})                 { println(logrus.FatalLevel, args...) }
func Fatalf(format string, args ...interface{})   { printf(logrus.FatalLevel, format, args...) }

func printf(level logrus.Level, format string, args ...interface{}) {
	NewEntry(RootLogger().Logger).withFileField().entry.Logf(level, format, args...)
}

func println(level logrus.Level, args ...interface{}) {
	NewEntry(RootLogger().Logger).withFileField().entry.Logln(level, args...)
}

func print(level logrus.Level, args ...interface{}) {
	NewEntry(RootLogger().Logger).withFileField().entry.Log(level, args...)
}
