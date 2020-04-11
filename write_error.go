package log

import (
	"fmt"
	"os"
)

// 缺省的写日志失败处理函数
func WriteError(data []byte, n int, err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "log write failed, data: ", string(data), ", n: ", n, ", err: ", err.Error())
	}
}
