package category_group_service

import (
	"database/sql"
	"fmt"
	"gofinn/internal/database"
	"gofinn/internal/models"
)

type SQLiteCategoryGroupService struct {
	db *sql.DB
}

func NewSQLiteCategoryGroupService() CategoryGroupService {
	return &SQLiteCategoryGroupService{
		db: database.GetDatabaseInstance(),
	}
}

func (s *SQLiteCategoryGroupService) ListCategoryGroupsForUser(
	ownerID int64,
) ([]models.CategoryGroup, error) {
	return s.getCategoryGroupsForUser(ownerID, 100)
}

func (s *SQLiteCategoryGroupService) CreateCategoryGroup(categoryGroup models.CategoryGroupFields) (models.CategoryGroup, error) {
	row := s.db.QueryRow(`
		INSERT INTO category_group (name, owner_id)
		VALUES (?, ?)
		RETURNING id, name, owner_id, created_at, updated_at, deleted_at
	`, categoryGroup.Name, categoryGroup.OwnerID)

	newCategoryGroup, err := scanIntoStruct(row)
	if err != nil {
		return models.CategoryGroup{}, err
	}
	return newCategoryGroup, nil
}

func (s *SQLiteCategoryGroupService) GetCategoryGroup(id int64) (models.CategoryGroup, error) {
	categoryGroups, err := s.getCategoryGroupsForUser(id, 1)
	if err != nil {
		return models.CategoryGroup{}, err
	}
	if len(categoryGroups) == 0 {
		return models.CategoryGroup{}, fmt.Errorf("category group with id %d not found", id)
	}
	return categoryGroups[0], nil
}

func (s *SQLiteCategoryGroupService) UpdateCategoryGroup(categoryGroup models.CategoryGroup) (models.CategoryGroup, error) {
	// Implement the method
	return categoryGroup, nil
}

func (s *SQLiteCategoryGroupService) DeleteCategoryGroup(id int64) error {
	// Implement the method
	return nil
}

func (s *SQLiteCategoryGroupService) getCategoryGroupsForUser(
	ownerID int64,
	limit int,
) ([]models.CategoryGroup, error) {
	var categoryGroups []models.CategoryGroup
	rows, err := s.db.Query(`
		SELECT
			id, name, owner_id,
			created_at, updated_at, deleted_at
		FROM category_group
		WHERE owner_id = ? AND deleted_at IS NULL
		LIMIT ?
	`, ownerID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		categoryGroup, err := scanIntoStruct(rows)
		if err == sql.ErrNoRows {
			break
		} else if err != nil {
			return nil, err
		}
		if categoryGroup.ID != 0 {
			categoryGroups = append(categoryGroups, categoryGroup)
		}
	}
	return categoryGroups, nil
}

func scanIntoStruct(row interface {
	Scan(dest ...interface{}) error
}) (models.CategoryGroup, error) {
	var categoryGroup models.CategoryGroup
	err := row.Scan(
		&categoryGroup.ID, &categoryGroup.Name, &categoryGroup.OwnerID,
		&categoryGroup.CreatedAt, &categoryGroup.UpdatedAt, &categoryGroup.DeletedAt,
	)
	return categoryGroup, err
}
