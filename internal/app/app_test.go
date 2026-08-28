package app

import (
	"path/filepath"
	"social-uploader/internal/store"
	"testing"
)

func TestServiceCreateChange(t *testing.T) {
	db, _ := store.Open(filepath.Join(t.TempDir(), "d"))
	defer db.Close()
	s := New(db)
	if _, e := s.Create("r", "b", "e", 3); e != nil {
		t.Fatal(e)
	}
	if e := s.Change("r", "changed"); e != nil {
		t.Fatal(e)
	}
}
