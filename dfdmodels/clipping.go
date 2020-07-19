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
	Ingredients []RecipeIngredient  `json:"ingredients"`
	Links       []Link              `json:"links"`
}
