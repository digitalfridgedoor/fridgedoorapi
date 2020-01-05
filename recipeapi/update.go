package recipeapi

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// Rename updates the name of the recipe
func Rename(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, name string) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	err = recipe.Rename(ctx, view.ID, recipeID, name)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, view, recipeID)
}

// UpdateMetadata updates the recipes metadata property
func UpdateMetadata(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, updates map[string]string) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	err = recipe.UpdateMetadata(ctx, view.ID, recipeID, updates)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, view, recipeID)
}
