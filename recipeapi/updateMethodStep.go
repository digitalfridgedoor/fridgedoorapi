package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddMethodStep adds a new method step to a recipe
func AddMethodStep(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID *primitive.ObjectID, action string) (*Recipe, error) {

	r, err := recipe.AddMethodStep(ctx, user.ViewID, recipeID, action)
	if err != nil {
		return nil, err
	}

	return mapToDto(r, user), nil
}

// UpdateMethodStep removes a method step
func UpdateMethodStep(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID *primitive.ObjectID, stepIdx int, updates map[string]string) (*Recipe, error) {

	r, err := recipe.UpdateMethodStepByIndex(ctx, user.ViewID, recipeID, stepIdx, updates)
	if err != nil {
		return nil, err
	}

	return mapToDto(r, user), nil
}

// RemoveMethodStep removes a method step
func RemoveMethodStep(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID *primitive.ObjectID, stepIdx int) (*Recipe, error) {

	r, err := recipe.RemoveMethodStepByIndex(ctx, user.ViewID, recipeID, stepIdx)
	if err != nil {
		return nil, err
	}

	return mapToDto(r, user), nil
}
