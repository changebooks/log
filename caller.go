package log

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// 调用函数的"文件名"和"代码行"
func Caller(skip int, short bool) string {
	_, file, line, ok := runtime.Caller(skip + 1)

	if ok {
		if short {
			file = filepath.Base(file)
		}
	} else {
		file = "???"
		line = 0
	}

	return fmt.Sprintf("%s:%d", file, line)
}
