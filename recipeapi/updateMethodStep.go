package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddMethodStep adds a new method step to a recipe
func AddMethodStep(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, action string) (*Recipe, error) {

	rID, err := primitive.ObjectIDFromHex(recipeID)
	if err != nil {
		return nil, errInvalidID
	}

	r, err := recipe.AddMethodStep(ctx, user.ViewID, &rID, action)
	if err != nil {
		return nil, err
	}

	return mapToDto(r, user), nil
}

// UpdateMethodStep removes a method step
func UpdateMethodStep(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, stepIdx int, updates map[string]string) (*Recipe, error) {

	rID, err := primitive.ObjectIDFromHex(recipeID)
	if err != nil {
		return nil, errInvalidID
	}

	r, err := recipe.UpdateMethodStepByIndex(ctx, user.ViewID, &rID, stepIdx, updates)
	if err != nil {
		return nil, err
	}

	return mapToDto(r, user), nil
}

// RemoveMethodStep removes a method step
func RemoveMethodStep(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, stepIdx int) (*Recipe, error) {

	rID, err := primitive.ObjectIDFromHex(recipeID)
	if err != nil {
		return nil, errInvalidID
	}

	r, err := recipe.RemoveMethodStepByIndex(ctx, user.ViewID, &rID, stepIdx)
	if err != nil {
		return nil, err
	}

	return mapToDto(r, user), nil
}
