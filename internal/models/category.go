package models

type CategoryFields struct {
	Name            string `json:"name" db:"name"`
	CategoryGroupID int64  `json:"category_group_id" db:"category_group_id"`
}

type Category struct {
	WithID
	CategoryFields
	Metadata
}

type CategoryFieldsWithID struct {
	WithID
	CategoryFields
}
