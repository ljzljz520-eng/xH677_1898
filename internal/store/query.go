package store

import (
	"go.etcd.io/bbolt"
	"social-uploader/internal/domain"
)

func (s *Store) UpdateRecord(r domain.Record) error { return s.SaveRecord(r) }
func (s *Store) DeleteRecord(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("records")).Delete([]byte(id)) })
}
