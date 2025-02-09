package models

type contextKey string

var ContextKeys = struct {
	Category   contextKey
	CategoryID contextKey
	User       contextKey
}{
	Category:   "category",
	CategoryID: "categoryID",
	User:       "user",
}
