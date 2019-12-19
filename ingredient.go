package fridgedoorapi

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoordatabase/ingredient"
)

// IngredientTreeNode represents a node in the ingredient tree
type IngredientTreeNode struct {
	ID    primitive.ObjectID    `json:"id" bson:"_id,omitempty"`
	Name  string                `json:"name"`
	Nodes []*IngredientTreeNode `json:"nodes"`
}

// CreateIngredient creates a new ingredient
func CreateIngredient(name string) (*ingredient.Ingredient, error) {
	return ingredient.Create(context.Background(), name)
}

// SearchIngredients retrieves the ingredients matching the query
func SearchIngredients(startsWith string) ([]*ingredient.Ingredient, error) {
	return ingredient.FindByName(context.Background(), startsWith)
}
