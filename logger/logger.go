package logger

type DebugLogger interface {
	Debug(args ...interface{})
	Debugln(args ...interface{})
	Debugf(format string, args ...interface{})
}

type InfoLogger interface {
	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(format string, args ...interface{})
}

type WarnLogger interface {
	Warn(args ...interface{})
	Warnln(args ...interface{})
	Warnf(format string, args ...interface{})
	Warning(args ...interface{})
	Warningln(args ...interface{})
	Warningf(format string, args ...interface{})
}

type ErrorLogger interface {
	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(format string, args ...interface{})
}

type FatalLogger interface {
	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(format string, args ...interface{})
}

type Logger interface {
	SetLevel(level string)
	DebugLogger
	InfoLogger
	WarnLogger
	ErrorLogger
	FatalLogger
}
