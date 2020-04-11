package log

import (
	"errors"
	"io"
)

// 队列节点
type QueueNode struct {
	writer io.WriteCloser
	data   []byte
}

func (x *QueueNode) Write() (n int, err error) {
	if x.writer == nil {
		return 0, errors.New("writer can't be nil")
	}

	if x.data == nil {
		return 0, errors.New("data can't be nil")
	}

	return x.writer.Write(x.data)
}
