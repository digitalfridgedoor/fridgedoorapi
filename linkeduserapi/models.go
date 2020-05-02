package linkeduserapi

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LinkedUser represents another user and their available recipes
type LinkedUser struct {
	ID       *primitive.ObjectID `json:"id"`
	Nickname string              `json:"nickname"`
	Recipes  []*Recipe           `json:"recipes"`
}

// Recipe is a view of another users recipe
type Recipe struct {
	ID    *primitive.ObjectID `json:"id"`
	Name  string              `json:"name"`
	Image bool                `json:"image"`
}
