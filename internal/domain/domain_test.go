package domain

import "testing"

func TestRecordLifecycle(t *testing.T) {
	r, e := NewRecord("r", "b", "e", 10)
	if e != nil {
		t.Fatal(e)
	}
	if e = r.Review(true, "ok"); e != nil {
		t.Fatal(e)
	}
	if e = r.Submit(); e != nil {
		t.Fatal(e)
	}
	if e = r.Archive(); e != nil {
		t.Fatal(e)
	}
	if !IsTerminal(r.Status) {
		t.Fatal(r.Status)
	}
}
