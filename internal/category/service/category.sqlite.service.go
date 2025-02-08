package category_service

import (
	"gofinn/internal/models"
)

type SQLiteCategoryService struct{}

func NewSQLiteCategoryService() CategoryService {
	return &SQLiteCategoryService{}
}

func (s *SQLiteCategoryService) ListCategories() ([]models.Category, error) {
	// Implement the method
	return []models.Category{}, nil
}

func (s *SQLiteCategoryService) CreateCategory(category models.Category) (models.Category, error) {
	// Implement the method
	return category, nil
}

func (s *SQLiteCategoryService) GetCategory(id int) (models.Category, error) {
	// Implement the method
	return models.Category{}, nil
}

func (s *SQLiteCategoryService) UpdateCategory(category models.Category) (models.Category, error) {
	// Implement the method
	return category, nil
}

func (s *SQLiteCategoryService) DeleteCategory(id int) error {
	// Implement the method
	return nil
}
