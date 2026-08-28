package validate

import (
	"social-uploader/internal/domain"
	"testing"
)

func TestBatchValidation(t *testing.T) {
	r, _ := domain.NewRecord("r", "b", "e", 4)
	if len(Record(r).Errors) != 0 {
		t.Fatal()
	}
	bad, _ := domain.NewRecord("x", "b", "e", -1)
	if len(Record(bad).Errors) == 0 {
		t.Fatal()
	}
}
