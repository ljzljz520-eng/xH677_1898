package report

import (
	"social-uploader/internal/domain"
	"social-uploader/internal/validate"
	"testing"
)

func TestReportDeterministic(t *testing.T) {
	r, _ := domain.NewRecord("z", "b", "e", 1)
	x := Build("b", []domain.Record{r}, []validate.Result{})
	if x.Text != "z" {
		t.Fatal(x.Text)
	}
}
