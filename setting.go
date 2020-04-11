package log

import (
	"runtime/debug"
	"strings"
	"sync"
)

// 配置
type Setting struct {
	enableCaller     bool                       // 是否报告调用函数的"文件名"和"代码行"
	callerShort      bool                       // 报告调用函数的"文件名"时，是否省略目录
	enableStacktrace bool                       // 发生错误时，是否追溯函数调用栈
	timeLayout       string                     // 时间模板，缺省：define.TimeLayout
	timeProcessor    func(layout string) string // 格式化时间函数
	formatter        func(s *Schema) []byte     // 格式化日志函数
	idGenerator      func() string              // id生成器
}

func (x *Setting) NowTime() string {
	return x.timeProcessor(x.timeLayout)
}

func (x *Setting) Caller(depth int) string {
	if x.enableCaller {
		return Caller(depth+1, x.callerShort)
	} else {
		return ""
	}
}

func (x *Setting) Stacktrace() string {
	if x.enableStacktrace {
		return string(debug.Stack())
	} else {
		return ""
	}
}

func (x *Setting) Format(s *Schema) []byte {
	return x.formatter(s)
}

func (x *Setting) NewId() string {
	return x.idGenerator()
}

func (x *Setting) GetEnableCaller() bool {
	return x.enableCaller
}

func (x *Setting) GetCallerShort() bool {
	return x.callerShort
}

func (x *Setting) GetEnableStacktrace() bool {
	return x.enableStacktrace
}

func (x *Setting) GetTimeLayout() string {
	return x.timeLayout
}

func (x *Setting) GetTimeProcessor() func(layout string) string {
	return x.timeProcessor
}

func (x *Setting) GetFormatter() func(s *Schema) []byte {
	return x.formatter
}

func (x *Setting) GetIdGenerator() func() string {
	return x.idGenerator
}

type SettingBuilder struct {
	mu               sync.Mutex // ensures atomic writes; protects the following fields
	enableCaller     bool
	callerShort      bool
	enableStacktrace bool
	timeLayout       string
	timeProcessor    func(layout string) string
	formatter        func(s *Schema) []byte
	idGenerator      func() string
}

func (x *SettingBuilder) Build() *Setting {
	timeLayout := x.timeLayout
	if timeLayout == "" {
		timeLayout = TimeLayout
	}

	timeProcessor := x.timeProcessor
	if timeProcessor == nil {
		timeProcessor = TimeProcessor
	}

	formatter := x.formatter
	if formatter == nil {
		formatter = JsonFormatter
	}

	idGenerator := x.idGenerator
	if idGenerator == nil {
		idGenerator = IdGenerator
	}

	return &Setting{
		enableCaller:     x.enableCaller,
		callerShort:      x.callerShort,
		enableStacktrace: x.enableStacktrace,
		timeLayout:       timeLayout,
		timeProcessor:    timeProcessor,
		formatter:        formatter,
		idGenerator:      idGenerator,
	}
}

func (x *SettingBuilder) SetEnableCaller(b bool) *SettingBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.enableCaller = b
	return x
}

func (x *SettingBuilder) SetCallerShort(b bool) *SettingBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.callerShort = b
	return x
}

func (x *SettingBuilder) SetEnableStacktrace(b bool) *SettingBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.enableStacktrace = b
	return x
}

func (x *SettingBuilder) SetTimeLayout(s string) *SettingBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.timeLayout = s
	return x
}

func (x *SettingBuilder) SetTimeProcessor(f func(layout string) string) *SettingBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.timeProcessor = f
	return x
}

func (x *SettingBuilder) SetFormatter(f func(s *Schema) []byte) *SettingBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.formatter = f
	return x
}

func (x *SettingBuilder) SetIdGenerator(f func() string) *SettingBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.idGenerator = f
	return x
}
