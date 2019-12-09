package appender

import (
	"github.com/hongwei-wu/log/internal/opts"
	"io"
	"sync/atomic"
)

const (
	Console = "consoleAppender"
)

type Appender interface {
	io.Writer
	SetLevel(level uint32)
	GetLevel() uint32
	GenPropOpt(prop string, value string) opts.Opt
	Apply(opts opts.Opts) error
}

func CreateAppender(name string) Appender {
	switch name {
	case Console:
		return NewConsoleAppender()
	default:
		return NewConsoleAppender()
	}
}

type LevelKeeper struct {
	level uint32
}

func (k *LevelKeeper) SetLevel(level uint32) { atomic.StoreUint32(&k.level, uint32(level)) }
func (k *LevelKeeper) GetLevel() uint32      { return atomic.LoadUint32(&k.level) }
