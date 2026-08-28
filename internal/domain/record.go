package domain

import (
	"errors"
	"fmt"
)

type Record struct {
	ID, BatchID, EmployeeID, Status, Remark string
	Amount                                  int64
}
type AuditEvent struct{ ID, RecordID, Action, Actor, Detail, At string }
type Workflow struct {
	ID, BatchID, State       string
	ErrorCount, SuccessCount int
}
type Attachment struct{ ID, BatchID, Name, Content string }

const (
	StatusDraft     = "draft"
	StatusApproved  = "approved"
	StatusFailed    = "failed"
	StatusSubmitted = "submitted"
	StatusArchived  = "archived"
)

func NewRecord(id, batch, employee string, amount int64) (Record, error) {
	if id == "" || batch == "" || employee == "" {
		return Record{}, errors.New("missing identity")
	}
	if amount <= 0 {
		return Record{}, errors.New("amount must be positive")
	}
	return Record{ID: id, BatchID: batch, EmployeeID: employee, Amount: amount, Status: StatusDraft, Remark: "registered"}, nil
}
func (r Record) Valid() bool {
	return r.ID != "" && r.BatchID != "" && r.EmployeeID != "" && r.Amount > 0
}
func (r Record) CanReview() bool { return r.Status == StatusDraft }
func (r *Record) Review(ok bool, note string) error {
	if !r.CanReview() {
		return errors.New("record not reviewable")
	}
	if !ok {
		r.Status = StatusFailed
		r.Remark = note
		return nil
	}
	r.Status = StatusApproved
	r.Remark = note
	return nil
}
func (r *Record) ChangeRemark(note string) error {
	if r.Status == StatusArchived {
		return errors.New("archived record")
	}
	if note == "" {
		return errors.New("empty remark")
	}
	r.Remark = note
	return nil
}
func (r *Record) Submit() error {
	if r.Status != StatusApproved {
		return fmt.Errorf("status %s cannot submit", r.Status)
	}
	r.Status = StatusSubmitted
	return nil
}
func (r *Record) Archive() error {
	if r.Status != StatusSubmitted {
		return errors.New("must submit before archive")
	}
	r.Status = StatusArchived
	return nil
}
func NewWorkflow(id, batch string) Workflow {
	return Workflow{ID: id, BatchID: batch, State: "created"}
}
func (w *Workflow) Count(success, failed int) {
	w.SuccessCount += success
	w.ErrorCount += failed
	if failed > 0 {
		w.State = "partial"
	} else {
		w.State = "submitted"
	}
}
func (w Workflow) Complete() bool { return w.State == "submitted" || w.State == "archived" }
