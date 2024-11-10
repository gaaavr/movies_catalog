package models

import (
	"context"
	"fmt"
)

const (
	addStateQuery = `INSERT INTO states (state_id, code, user_id)
					  VALUES ($1, $2, $3)`
	deleteStateQuery = `DELETE FROM states WHERE state_id = $1 AND code =$2 RETURNING user_id`
)

// AddState добавляет новое состояние
func (s *PG) AddState(ctx context.Context, stateID string, code, userID int64) error {
	_, err := s.db.ExecContext(ctx, addStateQuery, stateID, code, userID)
	if err != nil {
		return fmt.Errorf("[states_pg] failed to create state: %w", err)
	}

	return nil
}

// DeleteState удаляет состояние
func (s *PG) DeleteState(ctx context.Context, stateID string, code int64) (int64, error) {
	result := s.db.QueryRowContext(ctx, deleteStateQuery, stateID, code)
	var id int

	err := result.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("[states_pg] failed to delete state: %w", err)
	}

	return int64(id), nil
}
