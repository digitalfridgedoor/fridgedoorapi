package clippingapi

import "go.mongodb.org/mongo-driver/bson/primitive"

// ClippingDescription is a short view of a clipping
type ClippingDescription struct {
	ID       *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string              `json:"name"`
	RecipeID *primitive.ObjectID `json:"recipeID"`
}
