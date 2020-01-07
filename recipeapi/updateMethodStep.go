package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// AddMethodStep adds a new method step to a recipe
func AddMethodStep(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, action string) (*Recipe, error) {

	err := recipe.AddMethodStep(ctx, user.ViewID, recipeID, action)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, user, recipeID)
}

// UpdateMethodStep removes a method step
func UpdateMethodStep(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, stepIdx int, updates map[string]string) (*Recipe, error) {

	err := recipe.UpdateMethodStepByIndex(ctx, user.ViewID, recipeID, stepIdx, updates)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, user, recipeID)
}

// RemoveMethodStep removes a method step
func RemoveMethodStep(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, stepIdx int) (*Recipe, error) {

	err := recipe.RemoveMethodStepByIndex(ctx, user.ViewID, recipeID, stepIdx)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, user, recipeID)
}
