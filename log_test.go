package log

import "testing"

var (
	format = "%d %s %v"
	args   = []interface{}{1, "str_1", struct{}{}}
)

func TestDefaultLogger(t *testing.T) {
	t.Run("", func(t *testing.T) {
		SetLevel("debug")
		writeLog(t)
		SetLevel("info")
		writeLog(t)
		SetLevel("warn")
		writeLog(t)
		SetLevel("error")
		writeLog(t)
	})
}

func TestFatal(t *testing.T) {
	Fatal(args...)
}

func TestFatalln(t *testing.T) {
	Fatalln(args...)
}

func TestFatalf(t *testing.T) {
	Fatalf(format, args...)
}

func writeLog(t *testing.T) {
	t.Helper()

	Debug(args...)
	Debugln(args...)
	Debugf(format, args...)

	Info(args...)
	Infoln(args...)
	Infof(format, args...)

	Warn(args...)
	Warnln(args...)
	Warnf(format, args...)

	Warning(args...)
	Warningln(args...)
	Warningf(format, args...)

	Error(args...)
	Errorln(args...)
	Errorf(format, args...)
}
