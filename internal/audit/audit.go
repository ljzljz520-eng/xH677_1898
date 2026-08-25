package audit

import (
	"fmt"
	"social-uploader/internal/domain"
	"social-uploader/internal/store"
)

type Logger struct {
	next int
	db   *store.Store
}

func New(db *store.Store) *Logger { return &Logger{db: db} }
func (l *Logger) Event(recordID, action, actor, detail string) error {
	l.next++
	a := domain.AuditEvent{ID: fmt.Sprintf("audit-%04d", l.next), RecordID: recordID, Action: action, Actor: actor, Detail: detail, At: "deterministic"}
	return l.db.SaveAudit(a)
}
func (l *Logger) Register(id string) error {
	return l.Event(id, "register", "operator", "record registered")
}
func (l *Logger) Review(id, actor string, ok bool) error {
	action := "review-rejected"
	if ok {
		action = "review-approved"
	}
	return l.Event(id, action, actor, "review completed")
}
func (l *Logger) Submit(id string) error { return l.Event(id, "submit", "uploader", "submitted") }
