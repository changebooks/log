package log

import "sync"

type IdRegister struct {
	mu      sync.Mutex // ensures atomic writes; protects the following fields
	id      string     // 日志id
	traceId string     // 链路id
	bizId   string     // 业务id
}

func (x *IdRegister) GetId() string {
	return x.id
}

func (x *IdRegister) SetId(s string) string {
	x.mu.Lock()
	defer x.mu.Unlock()

	r := x.id
	x.id = s
	return r
}

func (x *IdRegister) GetTraceId() string {
	return x.traceId
}

func (x *IdRegister) SetTraceId(s string) {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.traceId = s
}

func (x *IdRegister) GetBizId() string {
	return x.bizId
}

func (x *IdRegister) SetBizId(s string) {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.bizId = s
}
