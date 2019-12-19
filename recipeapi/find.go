package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// FindOne finds a recipe by id
func FindOne(ctx context.Context, recipeID string) (*recipe.Recipe, error) {
	return recipe.FindOne(ctx, recipeID)
}
