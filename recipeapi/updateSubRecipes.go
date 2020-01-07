package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// AddSubRecipe adds a link between the recipe and the subrecipe
func AddSubRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, subRecipeID string) (*Recipe, error) {

	err := recipe.AddSubRecipe(ctx, user.ViewID, recipeID, subRecipeID)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, user, recipeID)
}

// RemoveSubRecipe the link between the recipe/subrecipe
func RemoveSubRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, subRecipeID string) (*Recipe, error) {

	err := recipe.RemoveSubRecipe(ctx, user.ViewID, recipeID, subRecipeID)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, user, recipeID)
}
