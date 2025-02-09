package models

type UserFields struct {
	Username       string `json:"username" db:"username"`
	HashedPassword string `json:"hashed_password" db:"hashed_password"`
	Currency       string `json:"currency" db:"currency"`
}

type User struct {
	WithID
	UserFields
	Metadata
}

type UserFieldsWithID struct {
	WithID
	UserFields
}
