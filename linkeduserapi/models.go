package linkeduserapi

import (
	"fridgedoorapi/search"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LinkedUser represents another user and their available recipes
type LinkedUser struct {
	ID       *primitive.ObjectID         `json:"id"`
	Nickname string                      `json:"nickname"`
	Recipes  []*search.RecipeDescription `json:"recipes"`
}
