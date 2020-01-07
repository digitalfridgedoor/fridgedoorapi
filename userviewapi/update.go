package userviewapi

import (
	"context"

	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/digitalfridgedoor/fridgedoorapi/apigateway"

	"github.com/digitalfridgedoor/fridgedoordatabase/userview"
)

// RemoveTag removes a tag from a recipe
func RemoveTag(ctx context.Context, request *apigateway.AuthenticatedUser, recipeID string, tag string) (*View, error) {
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
