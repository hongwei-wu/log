package appender

import (
	"encoding/binary"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	Black int32 = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const Clear = 0

const escape = "\x1b"

type ColorConsoleAppender struct {
	*os.File
	LevelKeeper
}

func NewColorConsoleAppender() *ColorConsoleAppender {
	a := &ColorConsoleAppender{File: os.Stdout}
	a.SetLevel(uint32(logrus.DebugLevel))
	return a
}

func (a *ColorConsoleAppender) Name() string {
	return "ColorConsole"
}

func (a *ColorConsoleAppender) Write(p []byte) (n int, err error) {
	level := binary.LittleEndian.Uint32(p)
	if uint32(a.GetLevel()) >= level {
		a.File.WriteString(fmt.Sprintf("%s[%dm", escape, levelColor(logrus.Level(level))))
		a.File.Write(p[4:])
		a.File.WriteString(fmt.Sprintf("%s[%dm", escape, Clear))
	}
	return len(p), nil
}

func levelColor(level logrus.Level) int32 {
	switch level {
	case logrus.DebugLevel:
		return Clear
	case logrus.InfoLevel:
		return Green
	case logrus.WarnLevel:
		return Yellow
	case logrus.ErrorLevel:
		return Red
	case logrus.FatalLevel:
		return White
	}
	return Clear
}
