package validate

import "social-uploader/internal/domain"

func HasDuplicate(records []domain.Record) bool {
	seen := map[string]bool{}
	for _, r := range records {
		if seen[r.EmployeeID] {
			return true
		}
		seen[r.EmployeeID] = true
	}
	return false
}
func Partition(results []Result) ([]domain.Record, []Result) {
	ok := []domain.Record{}
	bad := []Result{}
	for _, r := range results {
		if len(r.Errors) == 0 {
			ok = append(ok, r.Record)
		} else {
			bad = append(bad, r)
		}
	}
	return ok, bad
}
func TotalAmount(records []domain.Record) int64 {
	var n int64
	for _, r := range records {
		n += r.Amount
	}
	return n
}
