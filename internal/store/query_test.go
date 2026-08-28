package store

import (
	"path/filepath"
	"social-uploader/internal/domain"
	"testing"
)

func TestListAndDelete(t *testing.T) {
	s, e := Open(filepath.Join(t.TempDir(), "x"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r, _ := domain.NewRecord("a", "b", "e", 2)
	s.SaveRecord(r)
	xs, _ := s.ListRecords("b")
	if len(xs) != 1 {
		t.Fatal(len(xs))
	}
	if e = s.DeleteRecord("a"); e != nil {
		t.Fatal(e)
	}
}
