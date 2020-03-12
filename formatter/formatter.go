package formatter

import (
	"github.com/sirupsen/logrus"
)

type Formatter interface {
	Format(entry *logrus.Entry) ([]byte, error)
}

type Field struct {
	key   string
	field interface{}
}
