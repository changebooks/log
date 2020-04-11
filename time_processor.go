package log

import "time"

// 缺省的格式化时间函数，包含毫秒
func TimeProcessor(layout string) string {
	return time.Now().Format(layout + ".000")
}
