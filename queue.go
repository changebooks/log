package log

import (
	"errors"
	"io"
	"sync"
)

// 日志队列
type Queue struct {
	size uint64          // 队列-长度
	pipe chan *QueueNode // 队列
	pool *sync.Pool      // 节点复用池
}

func NewQueue(size uint64) *Queue {
	if size == 0 {
		size = QueueSize
	}

	return &Queue{
		size: size,
		pipe: make(chan *QueueNode, size),
		pool: &sync.Pool{
			New: func() interface{} {
				return &QueueNode{}
			},
		},
	}
}

func (x *Queue) Write(node *QueueNode) (n int, err error) {
	if node != nil {
		n, err = node.Write()
		x.PutPool(node)
		return
	} else {
		return 0, errors.New("node can't be nil")
	}
}

func (x *Queue) PutPipe(writer io.WriteCloser, data []byte) {
	node := x.GetPool()
	node.writer = writer
	node.data = data
	x.pipe <- node
}

// 不关闭
func (x *Queue) ClosePipe() {
	close(x.pipe)
}

func (x *Queue) PutPool(node *QueueNode) {
	if node != nil {
		x.pool.Put(node)
	}
}

func (x *Queue) GetPool() *QueueNode {
	return x.pool.Get().(*QueueNode)
}

func (x *Queue) GetSize() uint64 {
	return x.size
}
