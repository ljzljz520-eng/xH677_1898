package app

import (
	"fmt"
	"social-uploader/internal/domain"
	"social-uploader/internal/report"
)

func (s *Service) Create(id, batch, employee string, amount int64) (domain.Record, error) {
	r, e := domain.NewRecord(id, batch, employee, amount)
	if e != nil {
		return r, e
	}
	e = s.Register(r)
	return r, e
}
func (s *Service) Approve(id, actor string) error { return s.Review(id, actor, true, "approved") }
func (s *Service) Publish(r report.Report) string {
	return fmt.Sprintf("%s:%d", r.BatchID, len(r.Submitted))
}
func (s *Service) ValidateArchive(id string) error {
	r, e := s.DB.GetRecord(id)
	if e != nil {
		return e
	}
	if r.Status != domain.StatusSubmitted {
		return fmt.Errorf("not ready")
	}
	return s.Archive(id)
}
