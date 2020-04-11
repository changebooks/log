package log

import (
	"os"
	"testing"
)

func TestQueueSize(t *testing.T) {
	queue := NewQueue(1)
	got := queue.GetSize()
	var want uint64 = 1
	if got != 1 {
		t.Errorf("got %d; want %d", got, want)
	}

	queue2 := NewQueue(0)
	got2 := queue2.GetSize()
	want2 := QueueSize
	if got2 != want2 {
		t.Errorf("got %d; want %d", got2, want2)
	}
}

func TestQueueWrite(t *testing.T) {
	queue := NewQueue(0)

	_, got := queue.Write(nil)
	want := "node can't be nil"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	_, got2 := queue.Write(&QueueNode{})
	want2 := "writer can't be nil"
	if got2 == nil || got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	_, got3 := queue.Write(&QueueNode{writer: os.Stdout})
	want3 := "data can't be nil"
	if got3 == nil || got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}
}
