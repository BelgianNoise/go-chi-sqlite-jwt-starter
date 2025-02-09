package models

type CategoryFields struct {
	Name    string `json:"name" db:"name"`
	GroupID int64  `json:"group_id" db:"group_id"`
}

type Category struct {
	WithID
	WithOwnerID
	CategoryFields
	Metadata
}

type CategoryFieldsWithID struct {
	WithID
	CategoryFields
}
