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
	return &SQLiteCategoryService{
		db: database.GetDatabaseInstance(),
	}
}

func (s *SQLiteCategoryService) ListCategoriesForUser(
	ownerID int64,
) ([]models.Category, error) {
	return s.getCategoriesForUser(ownerID, 100)
}

func (s *SQLiteCategoryService) CreateCategory(category models.CategoryFields) (models.Category, error) {
	row := s.db.QueryRow(`
		INSERT INTO category (name, category_group_id)
		VALUES (?, ?)
		RETURNING id, name, category_group_id, created_at, updated_at, deleted_at
	`, category.Name, category.CategoryGroupID)

	newCategory, err := scanIntoStruct(row)
	if err != nil {
		return models.Category{}, err
	}
	return newCategory, nil
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
			c.id, c.name, c.category_group_id,
			c.created_at, c.updated_at, c.deleted_at
		FROM category as c INNER JOIN category_group as cg ON c.category_group_id = cg.id
		WHERE cg.owner_id = ? AND c.deleted_at IS NULL
		LIMIT ?
	`, ownerID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		category, err := scanIntoStruct(rows)
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

func scanIntoStruct(row interface {
	Scan(dest ...interface{}) error
}) (models.Category, error) {
	var category models.Category
	err := row.Scan(
		&category.ID, &category.Name, &category.CategoryGroupID,
		&category.CreatedAt, &category.UpdatedAt, &category.DeletedAt,
	)
	return category, err
}
