package recipeapi

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// AddMethodStep adds a new method step to a recipe
func AddMethodStep(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, action string) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	err = recipe.AddMethodStep(ctx, view.ID, recipeID, action)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, view, recipeID)
}

// UpdateMethodStep removes a method step
func UpdateMethodStep(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, stepIdx int, updates map[string]string) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	err = recipe.UpdateMethodStepByIndex(ctx, view.ID, recipeID, stepIdx, updates)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, view, recipeID)
}

// RemoveMethodStep removes a method step
func RemoveMethodStep(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, stepIdx int) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	err = recipe.RemoveMethodStepByIndex(ctx, view.ID, recipeID, stepIdx)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, view, recipeID)
}
