package domain

import "errors"

func Transition(r *Record, target string) error {
	switch target {
	case StatusApproved:
		return r.Review(true, "approved")
	case StatusSubmitted:
		return r.Submit()
	case StatusArchived:
		return r.Archive()
	case StatusFailed:
		r.Status = StatusFailed
		return nil
	default:
		return errors.New("unknown transition")
	}
}
func IsTerminal(status string) bool { return status == StatusArchived || status == StatusFailed }
func Clone(r Record) Record {
	return Record{ID: r.ID, BatchID: r.BatchID, EmployeeID: r.EmployeeID, Status: r.Status, Remark: r.Remark, Amount: r.Amount}
}
