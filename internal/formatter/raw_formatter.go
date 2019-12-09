package formatter

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/hongwei-wu/log/internal/opts"
	"github.com/sirupsen/logrus"
	"sort"
)

type RawFormatter struct {
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

	f.appendKeyValue(b, "time", entry.Time.Format("2006-01-02 15:04:05"))
	level := entry.Level.String()
	if level == "warning" {
		level = "warn"
	}
	f.appendKeyValue(b, "level", fmt.Sprintf("%-5s", level))

	/*
		if _, ok := entry.Data[fieldFile]; ok {
			f.appendKeyValue(b, fieldFile, entry.Data[fieldFile])
		}

		fields := make([]Field, 0, len(entry.Data))
		for k, f := range entry.Data {
			if k == fieldFile || k == FiledSkip {
				continue
			}
			fields = append(fields, KeyField{key: k, field: f})
		}
	*/

	fields := make([]Field, 0, len(entry.Data))
	for k, f := range entry.Data {
		fields = append(fields, Field{key: k, field: f})
	}
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].key < fields[j].key
	})

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
	if key == FieldFile {
		file := fmt.Sprint(value)
		f.appendValue(b, fmt.Sprintf("%-"+fmt.Sprintf("%d", 10)+"s", file))
	} else {
		f.appendValue(b, value)
	}
}

func (f *RawFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	b.WriteString(stringVal)
}

func (f *RawFormatter) GenPropOpt(prop string, value string) opts.Opt { return nil }
func (f *RawFormatter) Apply(opts opts.Opts) error                    { return nil }
