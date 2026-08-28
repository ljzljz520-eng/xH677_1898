package audit

import "testing"

func TestHelpers(t *testing.T) {
	if NormalizeActor(" Bob ") != "bob" {
		t.Fatal()
	}
	if EventKey("b", "x") != "b:x" {
		t.Fatal()
	}
	if Actions([]string{"x", "x"})["x"] != 2 {
		t.Fatal()
	}
}
