package category_service

import (
	"database/sql"
	"fmt"
	"gofinn/internal/database"
	"gofinn/internal/models"
)

type SQLiteCategoryService struct {
	db *sql.DB
}

func NewSQLiteCategoryService() CategoryService {
	db := database.GetDatabaseInstance()
	return &SQLiteCategoryService{db: db}
}

func (s *SQLiteCategoryService) ListCategoriesForUser(
	ownerID int64,
) ([]models.Category, error) {
	return s.getCategoriesForUser(ownerID, 100)
}

func (s *SQLiteCategoryService) CreateCategory(category models.Category) (models.Category, error) {
	// Implement the method
	return category, nil
}

func (s *SQLiteCategoryService) GetCategory(id int64) (models.Category, error) {
	categories, err := s.getCategoriesForUser(id, 1)
	if err != nil {
		return models.Category{}, err
	}
	if len(categories) == 0 {
		return models.Category{}, fmt.Errorf("category with id %d not found", id)
	}
	return categories[0], nil
}

func (s *SQLiteCategoryService) UpdateCategory(category models.Category) (models.Category, error) {
	// Implement the method
	return category, nil
}

func (s *SQLiteCategoryService) DeleteCategory(id int64) error {
	// Implement the method
	return nil
}

func (s *SQLiteCategoryService) getCategoriesForUser(
	ownerID int64,
	limit int,
) ([]models.Category, error) {
	var categories []models.Category
	rows, err := s.db.Query(`
		SELECT
			id, name, group_id, owner_id,
			created_at, updated_at, deleted_at
		FROM category
		WHERE owner_id = ? AND deleted_at IS NULL
		LIMIT ?
	`, ownerID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category)
		if err == sql.ErrNoRows {
			break
		} else if err != nil {
			return nil, err
		}
		if category.ID != 0 {
			categories = append(categories, category)
		}
	}
	return categories, nil
}
