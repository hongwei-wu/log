package formatter

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/hongwei-wu/log/field"
	"github.com/sirupsen/logrus"
	"sort"
	"sync/atomic"
)

type RawFormatter struct {
	maxFileWidth int32
}

func NewRawFormatter() *RawFormatter {
	return &RawFormatter{}
}

func (f *RawFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	if err := binary.Write(b, binary.LittleEndian, entry.Level); err != nil {
		return nil, err
	}

	f.appendValue(b, fmt.Sprintf("[%s][%s]", entry.Time.Format("2006-01-02 15:04:05"), f.levelAbbr(entry.Level.String())))

	if value, ok := entry.Data[field.File]; ok {
		file := fmt.Sprint(value)
		width := f.updateMaxFileWidth(file)
		f.appendValue(b, fmt.Sprintf("[%-"+fmt.Sprintf("%d", width)+"s]", file))
	}

	fields := make([]Field, 0, len(entry.Data))
	for k, f := range entry.Data {
		if k == field.File || k == field.Skip {
			continue
		}
		fields = append(fields, Field{key: k, field: f})
	}

	sort.Slice(fields, func(i, j int) bool { return fields[i].key < fields[j].key })
	for i := range fields {
		f.appendKeyValue(b, fields[i].key, fields[i].field)
	}

	if entry.Message != "" {
		f.appendKeyValue(b, "msg", entry.Message)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *RawFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString(key)
	b.WriteByte('=')
	f.appendValue(b, value)
}

func (f *RawFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	b.WriteString(stringVal)
}

func (f *RawFormatter) updateMaxFileWidth(file string) int32 {
	width := int32(len(file))
	max := atomic.LoadInt32(&f.maxFileWidth)
	if width <= max {
		width = max
	} else {
		atomic.StoreInt32(&f.maxFileWidth, width)
	}
	return width
}

func (f *RawFormatter) levelAbbr(level string) string {
	switch level {
	case "debug":
		return "D"
	case "info":
		return "I"
	case "warn":
		fallthrough
	case "warning":
		return "W"
	case "error":
		return "E"
	case "fatal":
		return "F"
	default:
		return "U"
	}
}
