package models

type contextKey string

var ContextKeys = struct {
	Category   contextKey
	CategoryID contextKey
}{
	Category:   "category",
	CategoryID: "categoryID",
}
