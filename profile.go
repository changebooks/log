package log

import (
	"errors"
	"os"
	"strconv"
	"sync"
)

type Profile struct {
	mu               sync.Mutex                                                  // ensures atomic writes; protects the following fields
	channel          string                                                      // 频道
	directory        string                                                      // 文件夹，必须可写权限
	level            int                                                         // 最小日志级别，小于此，不输出，缺省：define.Level
	timeLayout       string                                                      // 内容时间模板，缺省：define.TimeLayout
	fileTimeLayout   string                                                      // 文件名时间模板，缺省：define.FileTimeLayout
	fileExt          string                                                      // 文件名后缀，缺省：define.FileExt
	queueSize        uint64                                                      // 日志队列长度，缺省：define.QueueSize
	enableCaller     bool                                                        // 是否报告调用函数的"文件名"和"代码行"，缺省：define.EnableCaller
	callerShort      bool                                                        // 报告调用函数的"文件名"时，是否省略目录，缺省：define.CallerShort
	enableStacktrace bool                                                        // 发生错误时，是否追溯函数调用栈，缺省：define.EnableStacktrace
	perm             os.FileMode                                                 // 文件权限，缺省：define.FilePerm
	pathJoiner       func(directory string, channel string, level string) string // 拼接日志路径函数
	formatter        func(s *Schema) []byte                                      // 格式化日志函数
	idGenerator      func() string                                               // id生成器
	timeProcessor    func(layout string) string                                  // 格式化时间函数
}

func NewProfile(data map[string]string) (*Profile, error) {
	if data == nil {
		return nil, errors.New("data can't be nil")
	}

	level := -1
	if data[ProfileLevel] != "" {
		if l, err := strconv.ParseInt(data[ProfileLevel], 10, 32); err == nil {
			level = int(l)
		} else {
			return nil, err
		}
	}

	var queueSize uint64 = 0
	if data[ProfileQueueSize] != "" {
		if n, err := strconv.ParseUint(data[ProfileQueueSize], 10, 64); err == nil {
			queueSize = n
		} else {
			return nil, err
		}
	}

	enableCaller := EnableCaller
	if data[ProfileEnableCaller] != "" {
		if b, err := strconv.ParseBool(data[ProfileEnableCaller]); err == nil {
			enableCaller = b
		} else {
			return nil, err
		}
	}

	callerShort := CallerShort
	if data[ProfileCallerShort] != "" {
		if b, err := strconv.ParseBool(data[ProfileCallerShort]); err == nil {
			callerShort = b
		} else {
			return nil, err
		}
	}

	enableStacktrace := EnableStacktrace
	if data[ProfileEnableStacktrace] != "" {
		if b, err := strconv.ParseBool(data[ProfileEnableStacktrace]); err == nil {
			enableStacktrace = b
		} else {
			return nil, err
		}
	}

	var perm os.FileMode = 0
	if data[ProfilePerm] != "" {
		if p, err := strconv.ParseUint(data[ProfilePerm], 8, 32); err == nil {
			perm = os.FileMode(p)
		} else {
			return nil, err
		}
	}

	channel := data[ProfileChannel]
	directory := data[ProfileDirectory]
	timeLayout := data[ProfileTimeLayout]
	fileTimeLayout := data[ProfileFileTimeLayout]
	fileExt := data[ProfileFileExt]

	return &Profile{
		channel:          channel,
		directory:        directory,
		level:            level,
		timeLayout:       timeLayout,
		fileTimeLayout:   fileTimeLayout,
		fileExt:          fileExt,
		queueSize:        queueSize,
		enableCaller:     enableCaller,
		callerShort:      callerShort,
		enableStacktrace: enableStacktrace,
		perm:             perm,
	}, nil
}

func (x *Profile) GetChannel() string {
	return x.channel
}

func (x *Profile) GetDirectory() string {
	return x.directory
}

func (x *Profile) GetLevel() int {
	return x.level
}

func (x *Profile) GetTimeLayout() string {
	return x.timeLayout
}

func (x *Profile) GetFileTimeLayout() string {
	return x.fileTimeLayout
}

func (x *Profile) GetFileExt() string {
	return x.fileExt
}

func (x *Profile) GetQueueSize() uint64 {
	return x.queueSize
}

func (x *Profile) GetEnableCaller() bool {
	return x.enableCaller
}

func (x *Profile) GetCallerShort() bool {
	return x.callerShort
}

func (x *Profile) GetEnableStacktrace() bool {
	return x.enableStacktrace
}

func (x *Profile) GetPerm() os.FileMode {
	return x.perm
}

func (x *Profile) GetPathJoiner() func(directory string, channel string, level string) string {
	return x.pathJoiner
}

func (x *Profile) SetPathJoiner(f func(directory string, channel string, level string) string) *Profile {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.pathJoiner = f
	return x
}

func (x *Profile) GetFormatter() func(s *Schema) []byte {
	return x.formatter
}

func (x *Profile) SetFormatter(f func(s *Schema) []byte) *Profile {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.formatter = f
	return x
}

func (x *Profile) GetIdGenerator() func() string {
	return x.idGenerator
}

func (x *Profile) SetIdGenerator(f func() string) *Profile {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.idGenerator = f
	return x
}

func (x *Profile) GetTimeProcessor() func(layout string) string {
	return x.timeProcessor
}

func (x *Profile) SetTimeProcessor(f func(layout string) string) *Profile {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.timeProcessor = f
	return x
}
