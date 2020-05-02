package linkeduserapi

import (
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LinkedUser represents another user and their available recipes
type LinkedUser struct {
	ID       *primitive.ObjectID   `json:"id"`
	Nickname string                `json:"nickname"`
	Recipes  []*recipe.Description `json:"recipes"`
}
