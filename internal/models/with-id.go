package models

type WithID struct {
	ID string `json:"id" db:"id"`
}

type WithOwnerID struct {
	OwnerID string `json:"owner_id" db:"owner_id"`
}
