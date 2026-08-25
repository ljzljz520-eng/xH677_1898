package store

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"os"
	"social-uploader/internal/domain"
)

var buckets = []string{"records", "audits", "workflows", "attachments"}

type Store struct{ db *bbolt.DB }

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, n := range buckets {
			if _, x := tx.CreateBucketIfNotExists([]byte(n)); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return &Store{db: db}, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func encode(v any) []byte          { b, _ := json.Marshal(v); return b }
func decode(b []byte, v any) error { return json.Unmarshal(b, v) }
func (s *Store) SaveRecord(r domain.Record) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("records")).Put([]byte(r.ID), encode(r)) })
}
func (s *Store) GetRecord(id string) (domain.Record, error) {
	var r domain.Record
	e := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("records")).Get([]byte(id))
		if b == nil {
			return os.ErrNotExist
		}
		return decode(b, &r)
	})
	return r, e
}
func (s *Store) ListRecords(batch string) ([]domain.Record, error) {
	out := []domain.Record{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			var r domain.Record
			if err := decode(v, &r); err != nil {
				return err
			}
			if batch == "" || r.BatchID == batch {
				out = append(out, r)
			}
			return nil
		})
	})
	return out, e
}
func (s *Store) SaveWorkflow(w domain.Workflow) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("workflows")).Put([]byte(w.ID), encode(w)) })
}
func (s *Store) GetWorkflow(id string) (domain.Workflow, error) {
	var w domain.Workflow
	e := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("workflows")).Get([]byte(id))
		if b == nil {
			return os.ErrNotExist
		}
		return decode(b, &w)
	})
	return w, e
}
func (s *Store) SaveAudit(a domain.AuditEvent) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("audits")).Put([]byte(a.ID), encode(a)) })
}
func (s *Store) SaveAttachment(a domain.Attachment) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("attachments")).Put([]byte(a.ID), encode(a)) })
}
