package log

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"sync"
)

type Writer struct {
	queue   *QueueLoop             // 日志队列
	setting *Setting               // 配置
	level   int                    // 最小日志级别，小于此，不输出
	outs    map[int]io.WriteCloser // [ level => io.WriteCloser ]
}

func (x *Writer) Write(depth int, channel string, level int, name string, tag string, message interface{}, remark interface{}, err error, stacktrace string, idRegister *IdRegister) error {
	if x.Ignore(level) {
		return nil
	}

	out := x.GetOut(level)
	if out == nil {
		return fmt.Errorf("unsupported level %d, no out", level)
	}

	s, err := x.NewSchema(depth+1, channel, level, name, tag, message, remark, err, stacktrace, idRegister)
	if err != nil {
		return err
	}

	data := x.setting.Format(s)
	x.queue.Put(out, data)
	return nil
}

func (x *Writer) NewSchema(depth int, channel string, level int, name string, tag string, message interface{}, remark interface{}, err error, stacktrace string, idRegister *IdRegister) (*Schema, error) {
	l := GetLevel(level)
	if l == nil {
		return nil, fmt.Errorf("unsupported level %d, no level", level)
	}

	errStr := ""
	if err != nil {
		errStr = err.Error()
		if x.setting.GetEnableStacktrace() && stacktrace == "" {
			stacktrace = string(debug.Stack())
		}
	}

	var id string
	var parentId string
	var traceId string
	var bizId string
	if idRegister != nil {
		id = x.setting.NewId()
		parentId = idRegister.SetId(id)
		traceId = idRegister.GetTraceId()
		bizId = idRegister.GetBizId()
	}

	return &Schema{
		Id:         id,
		Name:       name,
		ParentId:   parentId,
		TraceId:    traceId,
		BizId:      bizId,
		Time:       x.setting.NowTime(),
		Channel:    channel,
		Level:      l.GetUpper(),
		Initial:    l.GetInitial(),
		Caller:     x.setting.Caller(depth + 1),
		ProcessId:  os.Getpid(),
		Tag:        tag,
		Message:    message,
		Remark:     remark,
		Error:      errStr,
		Stacktrace: stacktrace,
	}, nil
}

func (x *Writer) Ignore(level int) bool {
	return level < x.level
}

func (x *Writer) GetQueue() *QueueLoop {
	return x.queue
}

func (x *Writer) GetSetting() *Setting {
	return x.setting
}

func (x *Writer) GetLevel() int {
	return x.level
}

func (x *Writer) GetOut(level int) io.WriteCloser {
	return x.outs[level]
}

func (x *Writer) Close() []error {
	var r []error

	if x.outs != nil {
		for _, o := range x.outs {
			if o != nil {
				if err := o.Close(); err != nil {
					r = append(r, err)
				}
			}
		}
	}

	return r
}

type WriterBuilder struct {
	mu      sync.Mutex // ensures atomic writes; protects the following fields
	queue   *QueueLoop
	setting *Setting
	level   int
	outs    map[int]io.WriteCloser
}

func (x *WriterBuilder) Build() (*Writer, error) {
	if x.queue == nil {
		return nil, errors.New("queue can't be nil")
	}

	if x.setting == nil {
		return nil, errors.New("setting can't be nil")
	}

	if x.outs == nil {
		return nil, errors.New("outs can't be nil")
	}

	return &Writer{
		queue:   x.queue,
		setting: x.setting,
		level:   x.level,
		outs:    x.outs,
	}, nil
}

func (x *WriterBuilder) SetQueue(q *QueueLoop) *WriterBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.queue = q
	return x
}

func (x *WriterBuilder) SetSetting(s *Setting) *WriterBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.setting = s
	return x
}

func (x *WriterBuilder) SetLevel(l int) *WriterBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.level = l
	return x
}

func (x *WriterBuilder) SetOut(level int, out io.WriteCloser) *WriterBuilder {
	if out == nil {
		return x
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	if x.outs == nil {
		x.outs = make(map[int]io.WriteCloser)
	}

	x.outs[level] = out
	return x
}
