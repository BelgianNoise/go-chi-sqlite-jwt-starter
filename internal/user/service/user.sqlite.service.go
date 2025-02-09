package user_service

import (
	"database/sql"
	"fmt"
	"gofinn/internal/database"
	"gofinn/internal/models"
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
			id, username, hashed_password, currency
			created_at, updated_at, deleted_at
		FROM user
		WHERE deleted_at IS NULL
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user)
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
		INSERT INTO user (username, hashed_password, currency)
		VALUES (?, ?, ?)
		RETURNING id, username, hashed_password, currency, created_at, updated_at, deleted_at
	`, user.Username, user.HashedPassword, user.Currency)

	var newUser models.User
	err := row.Scan(&newUser)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func (s *SQLiteUserService) GetUser(id int64) (models.User, error) {
	var user models.User
	s.db.QueryRow(`
		SELECT
			id, username, hashed_password, currency
			created_at, updated_at, deleted_at
		FROM user
		WHERE id = ? AND deleted_at IS NULL
	`, id).Scan(&user)

	if user.ID == 0 {
		return models.User{}, fmt.Errorf("user with id %d not found", id)
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
