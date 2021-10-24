package userviewapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
)

func delete(ctx context.Context, username string) error {

	ok, coll := database.UserView(ctx)
	if !ok {
		return errNotConnected
	}

	view, err := GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	return coll.DeleteByID(ctx, view.ID)
}
