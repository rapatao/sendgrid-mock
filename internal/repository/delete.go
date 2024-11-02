package repository

import (
	"context"
	_ "embed"
)

var (
	//go:embed sql/delete.single.sql
	deleteSingleSQL string

	//go:embed sql/delete.all.sql
	deleteAllSQL string
)

func (s *Service) Delete(ctx context.Context, eventID string) error {
	_, err := s.conn.ExecContext(ctx, deleteSingleSQL, eventID)
	return err
}

func (s *Service) DeleteAll(ctx context.Context) error {
	_, err := s.conn.ExecContext(ctx, deleteAllSQL)
	return err
}
