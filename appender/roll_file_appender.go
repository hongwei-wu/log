package appender

import (
	"encoding/binary"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type RollFileAppender struct {
	LevelKeeper
	file    *os.File
	path    string
	maxFile int
	maxSize int
	curSize int
}

func NewRollFileAppender(path string, maxSize int, maxFile int) *RollFileAppender {
	sep := -1
	for i, c := range path {
		if os.IsPathSeparator(uint8(c)) {
			sep = i
		}
	}

	if sep != -1 {
		os.MkdirAll(path[0:sep], os.ModePerm)
	}

	roll := RollFileAppender{}
	roll.file = nil
	roll.path = path
	roll.maxSize = maxSize
	roll.maxFile = maxFile
	roll.curSize = 0
	roll.SetLevel(uint32(logrus.DebugLevel))

	return &roll
}

func (a *RollFileAppender) Name() string {
	return "RollFileAppender"
}

func (a *RollFileAppender) Write(p []byte) (int, error) {
	level := binary.LittleEndian.Uint32(p)
	if uint32(a.GetLevel()) >= level {
		a.writeBuf(p[4:])
	}
	return len(p), nil
}

func (a *RollFileAppender) writeBuf(buff []byte) {
	if !a.rollFile() {
		if a.file != nil {
			a.file.Close()
		}
		a.curSize = 0
		return
	}

	a.file.Write(buff)
	a.curSize += len(buff)
}

func (a *RollFileAppender) rollFile() bool {
	if a.file == nil {
		f, err := os.OpenFile(a.genFilePath(0), os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			return false
		}

		st, err := f.Stat()
		if err != nil {
			f.Close()
			return false
		}

		f.Seek(0, io.SeekEnd)
		a.file = f
		a.curSize = int(st.Size())
	}

	if a.curSize >= a.maxSize {
		a.file.Close()
		a.file = nil
		a.shiftFile()

		f, err := os.OpenFile(a.genFilePath(0), os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			return false
		}

		a.file = f
		a.curSize = 0
	}

	return true
}

func (a *RollFileAppender) shiftFile() bool {
	os.Remove(a.genFilePath(a.maxFile))
	for i := a.maxFile - 1; i >= 0; i-- {
		os.Rename(a.genFilePath(i), a.genFilePath(i+1))
	}
	return true
}

func (a *RollFileAppender) genFilePath(fileNum int) string {
	if fileNum == 0 {
		return a.path
	} else {
		return fmt.Sprintf("%s.%d", a.path, fileNum)
	}
}
