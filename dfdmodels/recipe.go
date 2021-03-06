package dfdmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Recipe represents a recipe
type Recipe struct {
	ID        *primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Name      string               `json:"name"`
	AddedOn   time.Time            `json:"addedOn"`
	AddedBy   primitive.ObjectID   `json:"addedBy"`
	Method    []MethodStep         `json:"method"`
	Recipes   []SubRecipe          `json:"recipes"`
	ParentIds []primitive.ObjectID `json:"parentIds"`
	Metadata  RecipeMetadata       `json:"metadata"`
}

// MethodStep is an instruction with a collection of ingredients
type MethodStep struct {
	Action      string             `json:"action"`
	Description string             `json:"description"`
	Time        string             `json:"time"`
	Ingredients []RecipeIngredient `json:"ingredients"`
}

// RecipeIngredient is the ingredient linked to each recipe
type RecipeIngredient struct {
	Name         string `json:"name"`
	Amount       string `json:"amount"`
	Preperation  string `json:"preperation"`
	IngredientID string `json:"ingredientId"`
}

// SubRecipe is a pointer to a recipe that makes up the main recipe
type SubRecipe struct {
	Name     string             `json:"name"`
	RecipeID primitive.ObjectID `json:"recipeId"`
}

// RecipeMetadata contains extra information about the recipe
type RecipeMetadata struct {
	Image       bool             `json:"image"`
	Tags        []string         `json:"tags"`
	RecipeLinks []RecipeLink     `json:"recipeLinks"`
	ViewableBy  RecipeViewableBy `json:"viewableBy"`
}

// RecipeLink is a link attached to a recipe
type RecipeLink struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// RecipeViewableBy describes who can view the recipe as well as the user
type RecipeViewableBy struct {
	Everyone bool                 `json:"everyone"`
	Users    []primitive.ObjectID `json:"users"`
}
