package userviewapi

import (
	"context"
	"fmt"

	"github.com/digitalfridgedoor/fridgedoordatabase/database"
)

func (editable *EditableView) save(ctx context.Context) bool {

	ok, coll := database.UserView(ctx)
	if !ok {
		return false
	}

	err := coll.UpdateByID(ctx, editable.db.ID, editable.db)
	if err != nil {
		fmt.Printf("Error saving userview: %v\n", err)
		return false
	}

	return true
}

func (editable *EditableView) saveAndGetDto(ctx context.Context) (*View, error) {
	ok := editable.save(ctx)
	if !ok {
		fmt.Printf("Did not save update, %v.\n", errNotConnected)
		return nil, errNotConnected
	}

	updated, err := findView(ctx, *editable.db.ID)
	if err != nil {
		return nil, err
	}

	return populateUserView(ctx, updated), nil
}
