package userviewapi

import (
	"context"
	"digitalfridgedoor/fridgedoordatabase/userview"

	"github.com/aws/aws-lambda-go/events"
)

// RemoveTag removes a tag from a recipe
func RemoveTag(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, tag string) (*View, error) {
	view, err := GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	err = userview.RemoveTag(ctx, view.ID.Hex(), tag)
	if err != nil {
		return nil, err
	}

	return GetUserViewByID(ctx, view.ID.Hex())
}
