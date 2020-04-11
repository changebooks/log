package log

import "testing"

func TestJsonFormatter(t *testing.T) {
	got := string(JsonFormatter(nil))
	want := "\n"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}

func TestNormalFormatter(t *testing.T) {
	got := string(NormalFormatter(nil))
	want := "\n"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}

func TestSimpleFormatter(t *testing.T) {
	got := string(SimpleFormatter(nil))
	want := "\n"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
