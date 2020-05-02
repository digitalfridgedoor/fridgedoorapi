package recipeapi

import (
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Recipe represents a recipe
type Recipe struct {
	ID          *primitive.ObjectID      `json:"id" bson:"_id,omitempty"`
	Name        string                   `json:"name"`
	OwnedByUser bool                     `json:"ownedByUser"`
	CanEdit     bool                     `json:"canEdit"`
	Method      []dfdmodels.MethodStep   `json:"method"`
	Recipes     []dfdmodels.SubRecipe    `json:"recipes"`
	ParentIds   []primitive.ObjectID     `json:"parentIds"`
	Metadata    dfdmodels.RecipeMetadata `json:"metadata"`
}
