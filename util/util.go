package util

import (
	"fmt"
	"path"
	"runtime"
)

func FileAndLine(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("%s:%-3d", path.Base(file), line)
}
