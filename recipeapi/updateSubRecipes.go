package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// AddSubRecipe adds a link between the recipe and the subrecipe
func AddSubRecipe(ctx context.Context, recipeID string, subRecipeID string) (*recipe.Recipe, error) {

	err := recipe.AddSubRecipe(ctx, recipeID, subRecipeID)
	if err != nil {
		return nil, err
	}

	return recipe.FindOne(ctx, recipeID)
}

// RemoveSubRecipe the link between the recipe/subrecipe
func RemoveSubRecipe(ctx context.Context, recipeID string, subRecipeID string) (*recipe.Recipe, error) {

	err := recipe.RemoveSubRecipe(ctx, recipeID, subRecipeID)
	if err != nil {
		return nil, err
	}

	return recipe.FindOne(ctx, recipeID)
}
