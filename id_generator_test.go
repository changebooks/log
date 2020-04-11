package log

import "testing"

func TestIdGenerator(t *testing.T) {
	got := len(IdGenerator())
	want := 32
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
