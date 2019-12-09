package log

import (
	"github.com/hongwei-wu/log/internal"
)

// levels
const (
	DEBUG = internal.DEBUG
	INFO  = internal.INFO
	WARN  = internal.WARN
	ERROR = internal.ERROR
	FATAL = internal.FATAL
)

func getRootLogger() Logger {
	return internal.GetDefaultLogger()
}

func SetLevel(level string) { getRootLogger().(*internal.Logger).SetLevel(level) }

func Debug(args ...interface{})                 { getRootLogger().Debug(args...) }
func Debugln(args ...interface{})               { getRootLogger().Debugln(args...) }
func Debugf(format string, args ...interface{}) { getRootLogger().Debugf(format, args...) }

func Info(args ...interface{})                 { getRootLogger().Info(args...) }
func Infoln(args ...interface{})               { getRootLogger().Infoln(args...) }
func Infof(format string, args ...interface{}) { getRootLogger().Infof(format, args...) }

func Warn(args ...interface{})                    { getRootLogger().Warn(args...) }
func Warnln(args ...interface{})                  { getRootLogger().Warnln(args...) }
func Warnf(format string, args ...interface{})    { getRootLogger().Warnf(format, args...) }
func Warning(args ...interface{})                 { getRootLogger().Warning(args...) }
func Warningln(args ...interface{})               { getRootLogger().Warningln(args...) }
func Warningf(format string, args ...interface{}) { getRootLogger().Warningf(format, args...) }

func Error(args ...interface{})                 { getRootLogger().Error(args...) }
func Errorln(args ...interface{})               { getRootLogger().Errorln(args...) }
func Errorf(format string, args ...interface{}) { getRootLogger().Errorf(format, args...) }

func Fatal(args ...interface{})                 { getRootLogger().Fatal(args...) }
func Fatalln(args ...interface{})               { getRootLogger().Fatalln(args...) }
func Fatalf(format string, args ...interface{}) { getRootLogger().Fatalf(format, args...) }
