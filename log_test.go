package log

import (
	"github.com/hongwei-wu/log/appender"
	"testing"
)

var (
	format = "%d %s %v"
	args   = []interface{}{1, "str_1", struct{}{}}
)

func TestLogger(t *testing.T) {

	t.Run("default appender", func(t *testing.T) {
		writeLogAllLevel(t)
		writeEntryLogAllLevel(t)
	})

	t.Run("color console appender", func(t *testing.T) {
		RootLogger().ResetAppender()
		RootLogger().AddAppender(appender.NewColorConsoleAppender())
		writeLogAllLevel(t)
		writeEntryLogAllLevel(t)
	})

	t.Run("role file appender", func(t *testing.T) {
		RootLogger().ResetAppender()
		RootLogger().AddAppender(appender.NewRollFileAppender("./log_dir/test.log", 10*K, 10))
		for i := 0; i < 1024; i++ {
			writeLogAllLevel(t)
			writeEntryLogAllLevel(t)
		}
	})

	t.Run("fatal", func(t *testing.T) {
		Fatal("Fatal")
	})
}

func writeLogAllLevel(t *testing.T) {
	SetLevel("debug")
	writeLog(t)
	SetLevel("info")
	writeLog(t)
	SetLevel("warn")
	writeLog(t)
	SetLevel("error")
	writeLog(t)
	SetLevel("fatal")
	writeLog(t)
}

func writeEntryLogAllLevel(t *testing.T) {
	SetLevel("debug")
	writeEntryLog(t)
	SetLevel("info")
	writeEntryLog(t)
	SetLevel("warn")
	writeEntryLog(t)
	SetLevel("error")
	writeEntryLog(t)
	SetLevel("fatal")
	writeEntryLog(t)
}

func writeEntryLog(t *testing.T) {
	t.Helper()
	WithField("field", "Debug").Debug("Debug")
	WithField("field", "Debugln").Debugln("Debugln")
	WithField("field", "Debugf").Debugf("%s", "Debugf")

	WithField("field", "Info").Info("Info")
	WithField("field", "Infoln").Infoln("Infoln")
	WithField("field", "Infof").Infof("%s", "Infof")

	WithField("field", "Warn").Warn("Warn")
	WithField("field", "Warnln").Warnln("Warnln")
	WithField("field", "Warnf").Warnf("%s", "Warnf")

	WithField("field", "Warning").Warning("Warning")
	WithField("field", "Warningln").Warningln("Warningln")
	WithField("field", "Warningf").Warningf("%s", "Warningf")

	WithField("field", "Error").Error("Error")
	WithField("field", "Errorln").Errorln("Errorln")
	WithField("field", "Errorf").Errorf("%s", "Errorf")
}

func writeLog(t *testing.T) {
	t.Helper()

	Debug("Debug")
	Debugln("Debugln")
	Debugf("%s", "Debugf")

	Info("Info")
	Infoln("Infoln")
	Infof("%s", "Infof")

	Warn("Warn")
	Warnln("Warnln")
	Warnf("%s", "Warnf")

	Warning("Warning")
	Warningln("Warningln")
	Warningf("%s", "Warningf")

	Error("Error")
	Errorln("Errorln")
	Errorf("%s", "Errorf")
}
