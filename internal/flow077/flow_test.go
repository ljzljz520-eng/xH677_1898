package flow077

import (
	"path/filepath"
	"social-uploader/internal/app"
	"social-uploader/internal/domain"
	"social-uploader/internal/store"
	"testing"
)

func setup(t *testing.T) (*Processor, func()) {
	db, e := store.Open(filepath.Join(t.TempDir(), "d"))
	if e != nil {
		t.Fatal(e)
	}
	return New(app.New(db)), func() { db.Close() }
}
func TestWorkflowImportReport(t *testing.T) {
	p, done := setup(t)
	defer done()
	good, _ := domain.NewRecord("g", "b", "e1", 5)
	bad, _ := domain.NewRecord("bad", "b", "e 2", 2)
	r, e := p.ImportBatch("b", []domain.Record{good, bad})
	if e != nil {
		t.Fatal(e)
	}
	if len(r.Submitted) != 1 || len(r.Errors) != 1 {
		t.Fatal(r)
	}
}
func TestWorkflowCreateReviewArchive(t *testing.T) {
	p, done := setup(t)
	defer done()
	r, _ := domain.NewRecord("r", "b", "e", 5)
	if e := p.Service.Register(r); e != nil {
		t.Fatal(e)
	}
	if e := p.Service.Approve("r", "auditor"); e != nil {
		t.Fatal(e)
	}
	if e := p.SubmitSuccessful("b", []string{"r"}); e != nil {
		t.Fatal(e)
	}
	if e := p.Service.Archive("r"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowSearchUpdatePublish(t *testing.T) {
	p, done := setup(t)
	defer done()
	r, _ := domain.NewRecord("r", "b", "e", 5)
	p.Service.Register(r)
	xs, e := p.Service.Search("b")
	if e != nil || len(xs) != 1 {
		t.Fatal(e, len(xs))
	}
	if e = p.Service.Change("r", "updated"); e != nil {
		t.Fatal(e)
	}
}
