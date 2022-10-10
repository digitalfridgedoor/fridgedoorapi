package recipeapi

import (
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ViewableRecipe struct {
	db   *dfdmodels.Recipe
	user *fridgedoorgateway.AuthenticatedUser
}

type EditableRecipe ViewableRecipe

type editableMethodStep struct {
	step *dfdmodels.MethodStep
}

// Recipe represents a recipe
type Recipe struct {
	ID          *primitive.ObjectID      `json:"id" bson:"_id,omitempty"`
	Name        string                   `json:"name"`
	Notes       string                   `json:"notes"`
	OwnedByUser bool                     `json:"ownedByUser"`
	CanEdit     bool                     `json:"canEdit"`
	Method      []dfdmodels.MethodStep   `json:"method"`
	Recipes     []dfdmodels.SubRecipe    `json:"recipes"`
	ParentIds   []primitive.ObjectID     `json:"parentIds"`
	Metadata    dfdmodels.RecipeMetadata `json:"metadata"`
}
