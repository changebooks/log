package log

import (
	"errors"
	"strings"
)

type Logger struct {
	stream *Stream
	name   string
	depth  int
}

func NewLogger(stream *Stream, name string, depth int) (*Logger, error) {
	if stream == nil {
		return nil, errors.New("stream can't be nil")
	}

	if name = strings.TrimSpace(name); name == "" {
		return nil, errors.New("name can't be empty")
	}

	if depth < 0 {
		depth = Depth
	}

	return &Logger{
		stream: stream,
		name:   name,
		depth:  depth,
	}, nil
}

func (x *Logger) V(tag string, message interface{}, remark interface{}, idRegister *IdRegister) {
	x.stream.Verbose(x.depth+1, x.name, tag, message, remark, idRegister)
}

func (x *Logger) D(tag string, message interface{}, remark interface{}, idRegister *IdRegister) {
	x.stream.Debug(x.depth+1, x.name, tag, message, remark, idRegister)
}

func (x *Logger) I(tag string, message interface{}, remark interface{}, idRegister *IdRegister) {
	x.stream.Info(x.depth+1, x.name, tag, message, remark, idRegister)
}

func (x *Logger) N(tag string, message interface{}, remark interface{}, idRegister *IdRegister) {
	x.stream.Notice(x.depth+1, x.name, tag, message, remark, idRegister)
}

func (x *Logger) W(tag string, message interface{}, remark interface{}, idRegister *IdRegister) {
	x.stream.Warning(x.depth+1, x.name, tag, message, remark, idRegister)
}

func (x *Logger) E(tag string, message interface{}, remark interface{}, err error, stacktrace string, idRegister *IdRegister) {
	x.stream.Error(x.depth+1, x.name, tag, message, remark, err, stacktrace, idRegister)
}

func (x *Logger) C(tag string, message interface{}, remark interface{}, err error, stacktrace string, idRegister *IdRegister) {
	x.stream.Critical(x.depth+1, x.name, tag, message, remark, err, stacktrace, idRegister)
}

func (x *Logger) A(tag string, message interface{}, remark interface{}, err error, stacktrace string, idRegister *IdRegister) {
	x.stream.Alert(x.depth+1, x.name, tag, message, remark, err, stacktrace, idRegister)
}

func (x *Logger) M(tag string, message interface{}, remark interface{}, err error, stacktrace string, idRegister *IdRegister) {
	x.stream.Emergency(x.depth+1, x.name, tag, message, remark, err, stacktrace, idRegister)
}

func (x *Logger) Close() []error {
	return x.stream.Close()
}

func (x *Logger) GetStream() *Stream {
	return x.stream
}

func (x *Logger) GetName() string {
	return x.name
}

func (x *Logger) GetDepth() int {
	return x.depth
}
