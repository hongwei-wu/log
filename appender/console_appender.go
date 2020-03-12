package appender

import (
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"os"
)

type ConsoleAppender struct {
	*os.File
	LevelKeeper
}

func NewConsoleAppender() *ConsoleAppender {
	a := &ConsoleAppender{File: os.Stdout}
	a.SetLevel(uint32(logrus.DebugLevel))
	return a
}

func (a *ConsoleAppender) Name() string {
	return "ConsoleAppender"
}

func (a *ConsoleAppender) Write(p []byte) (n int, err error) {
	level := binary.LittleEndian.Uint32(p)
	if uint32(a.GetLevel()) >= level {
		a.File.Write(p[4:])
	}
	return len(p), nil
}
