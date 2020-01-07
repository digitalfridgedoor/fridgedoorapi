package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/userview"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// Rename updates the name of the recipe
func Rename(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, name string) (*Recipe, error) {

	err := recipe.Rename(ctx, user.ViewID, recipeID, name)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, user, recipeID)
}

// UpdateMetadata updates the recipes metadata property
func UpdateMetadata(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, updates map[string]string) (*Recipe, error) {

	err := recipe.UpdateMetadata(ctx, user.ViewID, recipeID, updates)
	if err != nil {
		return nil, err
	}

	if update, ok := updates["tag_add"]; ok {
		userview.AddTag(ctx, user.ViewID.Hex(), update)
	}

	return findOneAndMap(ctx, user, recipeID)
}
