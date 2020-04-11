package log

import "testing"

func TestShort(t *testing.T) {
	got := Caller(0, true)
	want := "caller_test.go:6"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
