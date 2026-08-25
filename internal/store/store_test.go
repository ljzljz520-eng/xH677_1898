package store

import (
	"path/filepath"
	"social-uploader/internal/domain"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "data.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r, _ := domain.NewRecord("persist", "b", "e", 1)
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.GetRecord("persist")
	if e != nil || got.ID != "persist" {
		t.Fatalf("%v %#v", e, got)
	}
}
