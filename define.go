package log

import "os"

const (
	EOL                          = '\n'
	TimeLayout                   = "2006-01-02 15:04:05" // 缺省的Stream时间模板
	TimeLayout2                  = "15:04:05"            // 缺省的Terminal时间模板
	LabelError                   = "[error]"             // 错误日志标签
	LabelStacktrace              = "[stacktrace]"        // 堆栈日志标签
	FileExt                      = "log"
	FileTimeLayout               = "2006010215" // 文件名时间模板
	FilePerm         os.FileMode = 0766
	FileFlag                     = os.O_WRONLY | os.O_CREATE | os.O_APPEND
	QueueSize        uint64      = 1000  // 缺省的日志队列长度
	EnableCaller                 = true  // 是否报告调用函数的"文件名"和"代码行"
	CallerShort                  = false // 报告调用函数的"文件名"时，是否省略目录
	EnableStacktrace             = true  // 发生错误时，是否追溯函数调用栈
	Depth                        = 0
	IdSize                       = 16 // id长度
)

const (
	// Detailed information
	Verbose int = 1 + iota

	// Debug information
	Debug

	// Interesting events
	// Examples: User logs in, SQL logs.
	Info

	// Uncommon events
	Notice

	// Exceptional occurrences that are not errors
	// Examples: Use of deprecated APIs, poor use of an API, undesirable things that are not necessarily wrong.
	Warning

	// Runtime errors
	Error

	// Critical conditions
	// Example: Application component unavailable, unexpected exception.
	Critical

	// Action must be taken immediately
	// Example: Entire website down, database unavailable, etc.
	// This should trigger the SMS alerts and wake you up.
	Alert

	// Urgent alert
	Emergency

	// Ignore all
	Silent
)
