package domain

type BatchSummary struct {
	BatchID                            string
	Total, Approved, Submitted, Failed int
	Amount                             int64
}

func Summarize(records []Record, batch string) BatchSummary {
	s := BatchSummary{BatchID: batch}
	for _, r := range records {
		if r.BatchID != batch {
			continue
		}
		s.Total++
		s.Amount += r.Amount
		switch r.Status {
		case StatusApproved:
			s.Approved++
		case StatusSubmitted:
			s.Submitted++
		case StatusFailed:
			s.Failed++
		}
	}
	return s
}
func FilterBatch(records []Record, batch string) []Record {
	out := make([]Record, 0)
	for _, r := range records {
		if r.BatchID == batch {
			out = append(out, Clone(r))
		}
	}
	return out
}
