package models

type CategoryGroupFields struct {
	Name    string `json:"name" db:"name"`
	OwnerID int64  `json:"owner_id" db:"owner_id"`
}

type CategoryGroup struct {
	WithID
	CategoryGroupFields
	Metadata
}

type CategoryGroupFieldsWithID struct {
	WithID
	CategoryGroupFields
}
