package appender

import (
	"io"
	"sync/atomic"
)

type Appender interface {
	io.Writer
	Name() string
	SetLevel(level uint32)
	GetLevel() uint32
}

type LevelKeeper struct {
	level uint32
}

func (k *LevelKeeper) SetLevel(level uint32) { atomic.StoreUint32(&k.level, uint32(level)) }
func (k *LevelKeeper) GetLevel() uint32      { return atomic.LoadUint32(&k.level) }
