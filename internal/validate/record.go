package validate

import (
	"fmt"
	"social-uploader/internal/domain"
	"strings"
)

type Result struct {
	Record domain.Record
	Errors []string
}

func Record(r domain.Record) Result {
	out := Result{Record: r}
	if !r.Valid() {
		out.Errors = append(out.Errors, "identity or amount invalid")
	}
	if strings.Contains(r.EmployeeID, " ") {
		out.Errors = append(out.Errors, "employee id contains spaces")
	}
	if r.Amount > 100000000 {
		out.Errors = append(out.Errors, "amount exceeds limit")
	}
	return out
}
func Batch(records []domain.Record) []Result {
	out := make([]Result, 0, len(records))
	for _, r := range records {
		out = append(out, Record(r))
	}
	return out
}
func ErrorText(result Result) string {
	if len(result.Errors) == 0 {
		return ""
	}
	return fmt.Sprintf("%s: %s", result.Record.ID, strings.Join(result.Errors, "; "))
}
