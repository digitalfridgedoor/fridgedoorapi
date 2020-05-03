package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// CreateRecipe creates a new recipe with given name
func CreateRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, name string) (*Recipe, error) {

	recipe, err := recipe.Create(ctx, user.ViewID, name)

	return mapToDto(recipe, user), err
}
