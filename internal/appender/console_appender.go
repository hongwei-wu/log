package appender

import (
	"github.com/hongwei-wu/log/internal/opts"
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

func (a *ConsoleAppender) GenPropOpt(prop string, value string) opts.Opt { return nil }
func (a *ConsoleAppender) Apply(opts opts.Opts) error                    { return nil }
