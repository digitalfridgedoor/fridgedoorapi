package dfdmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Clipping represents a collection of links and lists of ingredients used while creating a recipe
type Clipping struct {
	ID          *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	AddedOn     time.Time           `json:"addedOn"`
	UserID      primitive.ObjectID  `json:"userID"`
	Name        string              `json:"name"`
	Ingredients []*RecipeIngredient `json:"ingredients"`
	Links       []ClippingLink      `json:"links"`
	RecipeID    *primitive.ObjectID `json:"recipeID"`
	Notes       string              `json:"notes"`
}

// ClippingLink is a link attached to a clipping
type ClippingLink struct {
	Name          string                `json:"name"`
	URL           string                `json:"url"`
	HasBeenParsed bool                  `json:"hasBeenParsed"`
	Ingredients   []*ClippingIngredient `json:"ingredients"`
	Notes         string                `json:"notes"`
}

// ClippingIngredient is an ingredient found in a link associated with a recipe
type ClippingIngredient struct {
	Name        string  `json:"name"`
	Amount      string  `json:"amount"`
	Preperation string  `json:"preperation"`
	Section     *string `json:"section"`
}
