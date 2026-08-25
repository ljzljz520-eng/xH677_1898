package flow077

import (
	"social-uploader/internal/domain"
	"sort"
)

func SortRecords(records []domain.Record) []domain.Record {
	out := append([]domain.Record{}, records...)
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}
func CountStatuses(records []domain.Record) map[string]int {
	m := map[string]int{}
	for _, r := range records {
		m[r.Status]++
	}
	return m
}
func Successful(records []domain.Record) []domain.Record {
	out := []domain.Record{}
	for _, r := range records {
		if r.Status == domain.StatusApproved {
			out = append(out, r)
		}
	}
	return out
}
