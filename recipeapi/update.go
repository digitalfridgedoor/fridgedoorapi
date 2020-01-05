package recipeapi

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"
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

// SetImageFlag indicates whether there is an image to get
func SetImageFlag(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, hasImage bool) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	err = recipe.SetImageFlag(ctx, view.ID, recipeID, hasImage)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, view, recipeID)
}

// AddTag adds a tag to a recipe
func AddTag(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, tag string) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	err = recipe.AddTag(ctx, view.ID, recipeID, tag)
	if err != nil {
		return nil, err
	}

	err = userview.AddTag(ctx, view.ID.Hex(), tag)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, view, recipeID)
}

// RemoveTag removes a tag from a recipe
func RemoveTag(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, tag string) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	err = recipe.RemoveTag(ctx, view.ID, recipeID, tag)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, view, recipeID)
}
