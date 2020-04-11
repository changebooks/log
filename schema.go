package log

type Schema struct {
	Id         string      `json:"id"`         // 日志id
	Name       string      `json:"name"`       // 日志名
	ParentId   string      `json:"parent"`     // 上级日志id
	TraceId    string      `json:"trace"`      // 链路id
	BizId      string      `json:"biz"`        // 业务id
	Time       string      `json:"time"`       // 时间
	Channel    string      `json:"channel"`    // 频道
	Level      string      `json:"level"`      // 级别
	Initial    string      `json:"initial"`    // 级别缩写
	Caller     string      `json:"caller"`     // 调用函数的"文件名"和"代码行"
	ProcessId  int         `json:"process"`    // 进程id
	Tag        string      `json:"tag"`        // 标签
	Message    interface{} `json:"message"`    // 内容
	Remark     interface{} `json:"remark"`     // 备注
	Error      string      `json:"error"`      // 错误消息
	Stacktrace string      `json:"stacktrace"` // 发生错误时，追溯函数调用栈
}
