package recipeapi

import (
	"fridgedoorapi/dfdmodels"
	"fridgedoorapi/fridgedoorgateway"

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
	OwnedByUser bool                     `json:"ownedByUser"`
	CanEdit     bool                     `json:"canEdit"`
	Method      []dfdmodels.MethodStep   `json:"method"`
	Recipes     []dfdmodels.SubRecipe    `json:"recipes"`
	ParentIds   []primitive.ObjectID     `json:"parentIds"`
	Metadata    dfdmodels.RecipeMetadata `json:"metadata"`
}
