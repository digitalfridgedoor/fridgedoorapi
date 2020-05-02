package recipeapi

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoordatabase/userview"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// Rename updates the name of the recipe
func Rename(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID *primitive.ObjectID, name string) (*Recipe, error) {

	r, err := recipe.Rename(ctx, user.ViewID, recipeID, name)
	if err != nil {
		return nil, err
	}

	return mapToDto(r, user), nil
}

// UpdateMetadata updates the recipes metadata property
func UpdateMetadata(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID *primitive.ObjectID, updates map[string]string) (*Recipe, error) {

	r, err := recipe.UpdateMetadata(ctx, user.ViewID, recipeID, updates)
	if err != nil {
		return nil, err
	}

	if update, ok := updates["tag_add"]; ok {
		userview.AddTag(ctx, &user.ViewID, update)
	}

	return mapToDto(r, user), nil
}
