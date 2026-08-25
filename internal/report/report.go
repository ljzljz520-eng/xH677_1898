package report

import (
	"social-uploader/internal/domain"
	"social-uploader/internal/validate"
	"sort"
	"strings"
)

type Report struct {
	BatchID   string
	Submitted []string
	Errors    []string
	Text      string
}

func Build(batch string, good []domain.Record, bad []validate.Result) Report {
	r := Report{BatchID: batch}
	for _, x := range good {
		r.Submitted = append(r.Submitted, x.ID)
	}
	for _, x := range bad {
		r.Errors = append(r.Errors, validate.ErrorText(x))
	}
	sort.Strings(r.Submitted)
	sort.Strings(r.Errors)
	r.Text = strings.Join(append(append([]string{}, r.Submitted...), r.Errors...), "\n")
	return r
}
func SuccessfulIDs(r Report) []string { return append([]string{}, r.Submitted...) }
func FailedDetails(r Report) []string { return append([]string{}, r.Errors...) }
