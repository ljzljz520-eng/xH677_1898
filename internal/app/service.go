package app

import (
	"errors"
	"social-uploader/internal/audit"
	"social-uploader/internal/domain"
	"social-uploader/internal/store"
)

type Service struct {
	DB    *store.Store
	Audit *audit.Logger
}

func New(db *store.Store) *Service { return &Service{DB: db, Audit: audit.New(db)} }
func (s *Service) Register(r domain.Record) error {
	if !r.Valid() {
		return errors.New("invalid record")
	}
	if e := s.DB.SaveRecord(r); e != nil {
		return e
	}
	return s.Audit.Register(r.ID)
}
func (s *Service) Review(id, actor string, ok bool, note string) error {
	r, e := s.DB.GetRecord(id)
	if e != nil {
		return e
	}
	if e = r.Review(ok, note); e != nil {
		return e
	}
	if e = s.DB.UpdateRecord(r); e != nil {
		return e
	}
	return s.Audit.Review(id, actor, ok)
}
func (s *Service) Change(id, note string) error {
	r, e := s.DB.GetRecord(id)
	if e != nil {
		return e
	}
	if e = r.ChangeRemark(note); e != nil {
		return e
	}
	return s.DB.UpdateRecord(r)
}
func (s *Service) Archive(id string) error {
	r, e := s.DB.GetRecord(id)
	if e != nil {
		return e
	}
	if e = r.Archive(); e != nil {
		return e
	}
	return s.DB.UpdateRecord(r)
}
func (s *Service) Search(batch string) ([]domain.Record, error) { return s.DB.ListRecords(batch) }
