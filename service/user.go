package service

import (
	"api/model"
	"context"
	"crypto/rand"
	"database/sql"
	"math/big"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, userName string, email string, password string) (*model.User, error) {
	const (
		insert  = `INSERT INTO users(user_name, email, password) VALUES(?, ?, ?)`
		confirm = `SELECT id, user_name, email, password, created_at, updated_at FROM users WHERE id = ?`
	)

	if _, err := s.db.PrepareContext(ctx, insert); err != nil {
		return nil, err
	}

	if _, err := s.db.PrepareContext(ctx, confirm); err != nil {
		return nil, err
	}

	result, err := s.db.ExecContext(ctx, insert, userName, email, password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, confirm, id)

	var user model.User
	row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	return &user, nil
}

func (s *UserService) LoginUser(ctx context.Context, email string, password string) (string, error) {
	const (
		read    = `SELECT id, email, password FROM users WHERE email = ? AND password = ?`
		insert  = `INSERT INTO tokens(user_id, token) VALUES(?, ?)`
		confirm = `SELECT token FROM tokens WHERE id = ?`
	)

	if _, err := s.db.PrepareContext(ctx, read); err != nil {
		return "", err
	}

	if _, err := s.db.PrepareContext(ctx, insert); err != nil {
		return "", err
	}

	if _, err := s.db.PrepareContext(ctx, confirm); err != nil {
		return "", err
	}

	row := s.db.QueryRowContext(ctx, read, email, password)

	var User model.User
	if err := row.Scan(&User.ID, &User.Email, &User.Password); err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}

	rand, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		return "", err
	}

	newToken := rand.String()

	result, err := s.db.ExecContext(ctx, insert, User.ID, newToken)
	if err != nil {
		return "", err
	}

	id, _ := result.LastInsertId()

	row = s.db.QueryRowContext(ctx, confirm, id)

	var insertedToken string
	row.Scan(&insertedToken)

	return newToken, nil
}

// for testing purpose
func (s *UserService) ReadUser(ctx context.Context, offsetID int64, limit int64) ([]*model.User, error) {
	const (
		read = `SELECT id, user_name, email, password, created_at, updated_at FROM users WHERE id >= ? ORDER BY id LIMIT ?`
	)

	_, err := s.db.PrepareContext(ctx, read)
	if err != nil {
		return nil, err
	}

	if limit == 0 {
		limit = 10
	}

	s.db.ExecContext(ctx, read, offsetID, limit) //不要？
	rows, err := s.db.QueryContext(ctx, read, offsetID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		users = append(users, &user)
	}
	return users, nil
}
