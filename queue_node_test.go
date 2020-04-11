package log

import (
	"os"
	"testing"
)

func TestQueueNodeWrite(t *testing.T) {
	queueNode := &QueueNode{}
	_, got := queueNode.Write()
	want := "writer can't be nil"
	if got == nil || got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	queueNode2 := &QueueNode{writer: os.Stdout}
	_, got2 := queueNode2.Write()
	want2 := "data can't be nil"
	if got2 == nil || got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}
}
