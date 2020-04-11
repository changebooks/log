package log

import (
	"io"
	"sync"
)

// 日志队列操作类
type QueueLoop struct {
	wg          sync.WaitGroup
	loopOnce    sync.Once
	queue       *Queue
	disposePipe chan int
}

func NewQueueLoop(size uint64) *QueueLoop {
	return &QueueLoop{
		queue:       NewQueue(size),
		disposePipe: make(chan int, 1),
	}
}

func (x *QueueLoop) LoopOnce() {
	x.loopOnce.Do(func() {
		x.wg.Add(1)
		go x.loopTake()
	})
}

func (x *QueueLoop) loopTake() {
	for {
		select {
		case node, ok := <-x.queue.pipe:
			if ok {
				_, _ = x.queue.Write(node)
			}
		case _, ok := <-x.disposePipe:
			if ok {
				x.wg.Done()
			}
		}
	}
}

func (x *QueueLoop) Put(writer io.WriteCloser, data []byte) {
	x.queue.PutPipe(writer, data)
	x.disposePipe <- 1
	x.wg.Wait()
	x.wg.Add(1)
}

func (x *QueueLoop) GetQueue() *Queue {
	return x.queue
}
