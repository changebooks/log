package log

import (
	"errors"
	"strings"
	"sync"
)

type Stream struct {
	channel string
	writer  *Writer
}

func (x *Stream) Verbose(depth int, name string, tag string, message interface{}, remark interface{}, idRegister *IdRegister) {
	_ = x.writer.Write(depth+1, x.channel, Verbose, name, tag, message, remark, nil, "", idRegister)
}

func (x *Stream) Debug(depth int, name string, tag string, message interface{}, remark interface{}, idRegister *IdRegister) {
	_ = x.writer.Write(depth+1, x.channel, Debug, name, tag, message, remark, nil, "", idRegister)
}

func (x *Stream) Info(depth int, name string, tag string, message interface{}, remark interface{}, idRegister *IdRegister) {
	_ = x.writer.Write(depth+1, x.channel, Info, name, tag, message, remark, nil, "", idRegister)
}

func (x *Stream) Notice(depth int, name string, tag string, message interface{}, remark interface{}, idRegister *IdRegister) {
	_ = x.writer.Write(depth+1, x.channel, Notice, name, tag, message, remark, nil, "", idRegister)
}

func (x *Stream) Warning(depth int, name string, tag string, message interface{}, remark interface{}, idRegister *IdRegister) {
	_ = x.writer.Write(depth+1, x.channel, Warning, name, tag, message, remark, nil, "", idRegister)
}

func (x *Stream) Error(depth int, name string, tag string, message interface{}, remark interface{}, err error, stacktrace string, idRegister *IdRegister) {
	_ = x.writer.Write(depth+1, x.channel, Error, name, tag, message, remark, err, stacktrace, idRegister)
}

func (x *Stream) Critical(depth int, name string, tag string, message interface{}, remark interface{}, err error, stacktrace string, idRegister *IdRegister) {
	_ = x.writer.Write(depth+1, x.channel, Critical, name, tag, message, remark, err, stacktrace, idRegister)
}

func (x *Stream) Alert(depth int, name string, tag string, message interface{}, remark interface{}, err error, stacktrace string, idRegister *IdRegister) {
	_ = x.writer.Write(depth+1, x.channel, Alert, name, tag, message, remark, err, stacktrace, idRegister)
}

func (x *Stream) Emergency(depth int, name string, tag string, message interface{}, remark interface{}, err error, stacktrace string, idRegister *IdRegister) {
	_ = x.writer.Write(depth+1, x.channel, Emergency, name, tag, message, remark, err, stacktrace, idRegister)
}

func (x *Stream) GetChannel() string {
	return x.channel
}

func (x *Stream) GetWriter() *Writer {
	return x.writer
}

func (x *Stream) Close() []error {
	return x.writer.Close()
}

type StreamBuilder struct {
	mu          sync.Mutex   // ensures atomic writes; protects the following fields
	channel     string       // 频道
	queue       *QueueLoop   // 日志队列
	setting     *Setting     // 配置
	fileBuilder *FileBuilder // 文件
	level       int          // 最小日志级别，小于此，不输出
}

func (x *StreamBuilder) Build() (*Stream, error) {
	if x.channel == "" {
		return nil, errors.New("channel can't be empty")
	}

	if x.queue == nil {
		return nil, errors.New("queue can't be nil")
	}

	if x.setting == nil {
		return nil, errors.New("setting can't be nil")
	}

	if x.fileBuilder == nil {
		return nil, errors.New("file builder can't be nil")
	}

	writerBuilder := &WriterBuilder{}
	writerBuilder.SetQueue(x.queue).SetSetting(x.setting).SetLevel(x.level)
	for _, l := range GetLevels() {
		if file, err := x.fileBuilder.SetLevel(l.GetLower()).Build(); err != nil {
			return nil, err
		} else {
			writerBuilder.SetOut(l.GetNum(), file)
		}
	}

	writer, err := writerBuilder.Build()
	if err != nil {
		return nil, err
	}

	return &Stream{
		channel: x.channel,
		writer:  writer,
	}, nil
}

func (x *StreamBuilder) SetChannel(s string) *StreamBuilder {
	s = strings.TrimSpace(s)

	x.mu.Lock()
	defer x.mu.Unlock()

	x.channel = s
	return x
}

func (x *StreamBuilder) SetQueue(q *QueueLoop) *StreamBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.queue = q
	return x
}

func (x *StreamBuilder) SetSetting(s *Setting) *StreamBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.setting = s
	return x
}

func (x *StreamBuilder) SetFileBuilder(b *FileBuilder) *StreamBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.fileBuilder = b
	return x
}

func (x *StreamBuilder) SetLevel(l int) *StreamBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.level = l
	return x
}
