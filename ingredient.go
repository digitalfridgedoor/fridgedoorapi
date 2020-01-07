package fridgedoorapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/ingredient"
)

// CreateIngredient creates a new ingredient
func CreateIngredient(name string) (*ingredient.Ingredient, error) {
	return ingredient.Create(context.Background(), name)
}

// SearchIngredients retrieves the ingredients matching the query
func SearchIngredients(startsWith string) ([]*ingredient.Ingredient, error) {
	return ingredient.FindByName(context.Background(), startsWith)
}
