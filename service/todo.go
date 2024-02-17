package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"test/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
// dbをメンバに持つTODOServiceを新しく作って返しているんだね
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
// 引数はcontext, subject, descriptionを受け取っているね
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	s.db.PrepareContext(ctx, insert)
	s.db.PrepareContext(ctx, confirm)

	result, err := s.db.ExecContext(ctx, insert, subject, description)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()

	row := s.db.QueryRowContext(ctx, confirm, id)

	var todo model.TODO
	row.Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)

	todo.ID = id

	if todo.Subject == "" {
		return &todo, nil
	}

	if todo.Description == "" {
		return &todo, nil
	}

	return &todo, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	s.db.PrepareContext(ctx, read)
	s.db.PrepareContext(ctx, readWithID)

	todos := []*model.TODO{}
	var rows *sql.Rows
	var err error

	//prev_id指定なし
	if prevID == 0 {
		if size == 0 {
			s.db.ExecContext(ctx, read, 100)
			rows, err = s.db.QueryContext(ctx, read, 100)
			if err != nil {
				return nil, err
			}
		} else {
			s.db.ExecContext(ctx, read, size)
			rows, err = s.db.QueryContext(ctx, read, size)
			if err != nil {
				return nil, err
			}
		}
	} else {
		if _, err := s.db.ExecContext(ctx, readWithID, prevID, size); err != nil {
			return nil, err
		}

		rows, err = s.db.QueryContext(ctx, readWithID, prevID, size)
		if err != nil {
			return nil, err
		}
	}

	for rows.Next() {
		var todo model.TODO
		rows.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		todos = append(todos, &todo)
	}

	return todos, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	s.db.PrepareContext(ctx, update)
	s.db.PrepareContext(ctx, confirm)

	result, err := s.db.ExecContext(ctx, update, subject, description, id)
	if err != nil {
		return nil, err
	}

	if count, _ := result.RowsAffected(); count == 0 {
		return nil, &model.ErrNotFound{}
	}

	row := s.db.QueryRowContext(ctx, confirm, id)

	var todo model.TODO
	row.Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)

	todo.ID = id

	return &todo, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	query := fmt.Sprintf(deleteFmt, strings.Repeat(",?", len(ids)-1))

	s.db.PrepareContext(ctx, deleteFmt)

	idArgs := []interface{}{}

	for _, id := range ids {
		idArgs = append(idArgs, id)
	}

	result, err := s.db.ExecContext(ctx, query, idArgs...)
	if err != nil {
		return err
	}

	count, _ := result.RowsAffected()

	if count == 0 {
		return &model.ErrNotFound{}
	}

	return nil
}
