package user_service

import (
	"database/sql"
	"fmt"
	"go-chi-sqlite-jwt-starter/internal/database"
	"go-chi-sqlite-jwt-starter/internal/models"
)

type SQLiteUserService struct {
	db *sql.DB
}

func NewSQLiteUserService() UserService {
	return &SQLiteUserService{
		db: database.GetDatabaseInstance(),
	}
}

func (s *SQLiteUserService) ListUsers() ([]models.User, error) {
	var users []models.User
	rows, err := s.db.Query(`
		SELECT
			id, username, hashed_password, currency, role,
			created_at, updated_at, deleted_at
		FROM user
		WHERE deleted_at IS NULL
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user, err := scanIntoStruct(rows)
		if err == sql.ErrNoRows {
			break
		} else if err != nil {
			return nil, err
		}
		if user.ID != 0 {
			users = append(users, user)
		}
	}
	return users, nil
}

func (s *SQLiteUserService) CreateUser(user models.UserFields) (models.User, error) {
	row := s.db.QueryRow(`
		INSERT INTO user (username, hashed_password)
		VALUES (?, ?, ?)
		RETURNING id, username, hashed_password, role, created_at, updated_at, deleted_at
	`, user.Username, user.HashedPassword)
	newUser, err := scanIntoStruct(row)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func (s *SQLiteUserService) GetUser(id int64) (models.User, error) {
	row := s.db.QueryRow(`
		SELECT
			id, username, hashed_password, role,
			created_at, updated_at, deleted_at
		FROM user
		WHERE id = ? AND deleted_at IS NULL
	`, id)
	user, err := scanIntoStruct(row)
	if err != nil {
		return models.User{}, err
	} else if user.ID == 0 {
		return models.User{}, fmt.Errorf("user with id %d not found", id)
	}

	return user, nil
}

func (s *SQLiteUserService) GetUserByUsername(username string) (models.User, error) {
	row := s.db.QueryRow(`
		SELECT
			id, username, hashed_password, role,
			created_at, updated_at, deleted_at
		FROM user
		WHERE username = ? AND deleted_at IS NULL
	`, username)
	user, err := scanIntoStruct(row)
	if err != nil {
		return models.User{}, err
	} else if user.ID == 0 {
		return models.User{}, fmt.Errorf("user with username %q not found", username)
	}

	return user, nil
}

func (s *SQLiteUserService) UpdateUser(user models.User) (models.User, error) {
	// Implement the method
	return user, nil
}

func (s *SQLiteUserService) DeleteUser(id int64) error {
	// Implement the method
	return nil
}

func scanIntoStruct(row interface {
	Scan(dest ...interface{}) error
}) (models.User, error) {
	var user models.User
	err := row.Scan(
		&user.ID, &user.Username, &user.HashedPassword, &user.Role,
		&user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)
	return user, err
}
