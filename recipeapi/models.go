package recipeapi

import (
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Recipe represents a recipe
type Recipe struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name        string               `json:"name"`
	OwnedByUser bool                 `json:"ownedByUser"`
	CanEdit     bool                 `json:"canEdit"`
	Method      []recipe.MethodStep  `json:"method"`
	Recipes     []recipe.SubRecipe   `json:"recipes"`
	ParentIds   []primitive.ObjectID `json:"parentIds"`
	Metadata    recipe.Metadata      `json:"metadata"`
}
