package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// AddMethodStep adds a new method step to a recipe
func AddMethodStep(ctx context.Context, recipeID string, action string) (*recipe.Recipe, error) {

	err := recipe.AddMethodStep(ctx, recipeID, action)
	if err != nil {
		return nil, err
	}

	return recipe.FindOne(ctx, recipeID)
}

// UpdateMethodStep removes a method step
func UpdateMethodStep(ctx context.Context, recipeID string, stepIdx int, updates map[string]string) (*recipe.Recipe, error) {

	err := recipe.UpdateMethodStepByIndex(ctx, recipeID, stepIdx, updates)
	if err != nil {
		return nil, err
	}

	return recipe.FindOne(ctx, recipeID)
}

// RemoveMethodStep removes a method step
func RemoveMethodStep(ctx context.Context, recipeID string, stepIdx int) (*recipe.Recipe, error) {

	err := recipe.RemoveMethodStepByIndex(ctx, recipeID, stepIdx)
	if err != nil {
		return nil, err
	}

	return recipe.FindOne(ctx, recipeID)
}
