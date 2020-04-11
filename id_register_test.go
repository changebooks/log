package log

import "testing"

func TestIdRegisterSetId(t *testing.T) {
	idRegister := &IdRegister{}

	got := idRegister.SetId("abc")
	want := ""
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := idRegister.SetId("def")
	want2 := "abc"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := idRegister.SetId("")
	want3 := "def"
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}
}
