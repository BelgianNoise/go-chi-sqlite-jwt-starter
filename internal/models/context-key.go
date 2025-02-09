package models

type contextKey string

var ContextKeys = struct {
	Category        contextKey
	CategoryID      contextKey
	CategoryGroup   contextKey
	CategoryGroupID contextKey
	User            contextKey
}{
	Category:        "category",
	CategoryID:      "categoryID",
	CategoryGroup:   "categoryGroup",
	CategoryGroupID: "categoryGroupID",
	User:            "user",
}
