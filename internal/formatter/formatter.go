package formatter

import (
	"github.com/hongwei-wu/log/internal/opts"
	"github.com/sirupsen/logrus"
)

type Formatter interface {
	Format(entry *logrus.Entry) ([]byte, error)
	GenPropOpt(prop string, value string) opts.Opt
	Apply(opts opts.Opts) error
}

type Field struct {
	key   string
	field interface{}
}

const (
	Raw = "rawFormatter"
)

func CreateFormatter(name string) Formatter {
	switch name {
	case Raw:
		return NewRawFormatter()
	default:
		return NewRawFormatter()
	}
}
