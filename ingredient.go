package fridgedoorapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/ingredient"
)

// SearchIngredients retrieves the ingredients matching the query
func SearchIngredients(startsWith string) ([]*ingredient.Ingredient, error) {
	i, err := Ingredient()
	if err != nil {
		return nil, err
	}

	return i.FindByName(context.Background(), startsWith)
}
