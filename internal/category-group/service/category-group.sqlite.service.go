package category_group_service

import (
	"database/sql"
	"fmt"
	"gofinn/internal/database"
	"gofinn/internal/models"
	"strings"
)

type SQLiteCategoryGroupService struct {
	db      *sql.DB
	columns []string
}

func NewSQLiteCategoryGroupService() CategoryGroupService {
	return &SQLiteCategoryGroupService{
		db: database.GetDatabaseInstance(),
		columns: []string{
			"id", "name", "owner_id",
			"created_at", "updated_at", "deleted_at",
		},
	}
}

func (s *SQLiteCategoryGroupService) ListCategoryGroupsForUser(
	ownerID int64,
) ([]models.CategoryGroup, error) {
	var categoryGroups []models.CategoryGroup
	rows, err := s.db.Query(`
		SELECT `+strings.Join(s.columns, ", ")+`
		FROM category_group
		WHERE owner_id = ? AND deleted_at IS NULL
	`, ownerID)
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

func (s *SQLiteCategoryGroupService) CreateCategoryGroup(categoryGroup models.CategoryGroupFields) (models.CategoryGroup, error) {
	row := s.db.QueryRow(`
		INSERT INTO category_group (name, owner_id)
		VALUES (?, ?)
		RETURNING `+strings.Join(s.columns, ", ")+`
	`, categoryGroup.Name, categoryGroup.OwnerID)

	newCategoryGroup, err := scanIntoStruct(row)
	if err != nil {
		return models.CategoryGroup{}, err
	}
	return newCategoryGroup, nil
}

func (s *SQLiteCategoryGroupService) GetCategoryGroupForUser(
	id int64,
	ownerID int64,
) (models.CategoryGroup, error) {
	row := s.db.QueryRow(`
		SELECT `+strings.Join(s.columns, ", ")+`
		FROM category_group
		WHERE id = ? AND owner_id = ? AND deleted_at IS NULL
	`, id, ownerID)
	categoryGroup, err := scanIntoStruct(row)
	if err == sql.ErrNoRows {
		return models.CategoryGroup{}, fmt.Errorf("Category group with ID %d not found", id)
	}
	return categoryGroup, err
}

func (s *SQLiteCategoryGroupService) GetCategoryGroup(id int64) (models.CategoryGroup, error) {
	row := s.db.QueryRow(`
		SELECT `+strings.Join(s.columns, ", ")+`
		FROM category_group
		WHERE id = ?
	`, id)
	categoryGroup, err := scanIntoStruct(row)
	if err == sql.ErrNoRows {
		return models.CategoryGroup{}, fmt.Errorf("Category group not found")
	}
	return categoryGroup, err
}

func (s *SQLiteCategoryGroupService) UpdateCategoryGroup(categoryGroup models.CategoryGroup) (models.CategoryGroup, error) {
	// Implement the method
	return categoryGroup, nil
}

func (s *SQLiteCategoryGroupService) DeleteCategoryGroup(id int64) error {
	// Implement the method
	return nil
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
