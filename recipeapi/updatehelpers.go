package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"
)

func findOneAndMap(ctx context.Context, view *userview.View, recipeID string) (*Recipe, error) {

	recipe, err := recipe.FindOne(ctx, recipeID)
	if err != nil {
		return nil, err
	}

	return mapToDto(recipe, view), nil
}
