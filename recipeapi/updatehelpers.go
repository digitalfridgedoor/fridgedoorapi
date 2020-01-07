package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

func findOneAndMap(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string) (*Recipe, error) {

	recipe, err := recipe.FindOne(ctx, recipeID)
	if err != nil {
		return nil, err
	}

	return mapToDto(recipe, user), nil
}
