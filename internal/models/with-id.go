package models

type WithID struct {
	ID int64 `json:"id" db:"id"`
}

type WithOwnerID struct {
	OwnerID int64 `json:"owner_id" db:"owner_id"`
}
