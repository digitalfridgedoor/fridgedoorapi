package userviewapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"github.com/digitalfridgedoor/fridgedoordatabase/userview"
)

// RemoveTag removes a tag from a recipe
func RemoveTag(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, tag string) (*View, error) {

	err = userview.RemoveTag(ctx, view.ID.Hex(), tag)
	if err != nil {
		return nil, err
	}

	return GetUserViewByID(ctx, view.ID.Hex())
}
