package flow077

import (
	"social-uploader/internal/domain"
	"testing"
)

func Test677BusinessRegression(t *testing.T) {
	p, done := setup(t)
	defer done()
	good, _ := domain.NewRecord("good", "batch-2", "E1", 10)
	bad, _ := domain.NewRecord("failed", "batch-2", "E 2", 10)
	if _, e := p.ImportBatch("batch-2", []domain.Record{good, bad}); e != nil {
		t.Fatal(e)
	}
	r, e := p.Service.DB.GetRecord("failed")
	if e != nil {
		t.Fatal(e)
	}
	if r.Status != domain.StatusFailed {
		t.Fatalf("failed item status=%s", r.Status)
	}
}
