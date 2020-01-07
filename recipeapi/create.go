package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"
)

// CreateRecipe creates a new recipe with given name
func CreateRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, collectionName string, name string) (*recipe.Recipe, error) {

	recipe, err := recipe.Create(ctx, user.ViewID, name)
	if err != nil {
		return nil, err
	}

	err = userview.AddRecipe(ctx, user.ViewID.Hex(), collectionName, recipe.ID)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}
