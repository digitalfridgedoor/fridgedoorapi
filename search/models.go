package search

import "go.mongodb.org/mongo-driver/bson/primitive"

// RecipeDescription is a short view of the recipe
type RecipeDescription struct {
	ID    *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string              `json:"name"`
	Image bool                `json:"image"`
}

// Result is a short view of the recipe
type Result struct {
	RecipeID   *primitive.ObjectID `json:"recipeId"`
	ClippingID *primitive.ObjectID `json:"clippingId"`
	Name       string              `json:"name"`
}
