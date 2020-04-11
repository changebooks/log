package log

import (
	"fmt"
	"path/filepath"
	"time"
)

// 缺省的拼接日志路径函数
func PathJoiner(directory string, channel string, level string) string {
	return filepath.Join(directory, FileNameJoiner(channel, level))
}

// 缺省的拼接日志文件名函数，channel.level-年月日时.log
func FileNameJoiner(channel string, level string) string {
	return fmt.Sprintf("%s.%s-%s.%s", channel, level, time.Now().Format(FileTimeLayout), FileExt)
}
