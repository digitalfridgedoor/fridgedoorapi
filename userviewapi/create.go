package userviewapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoordatabase/database"
)

// Create creates a new userview for a user
func Create(ctx context.Context, username string) (*dfdmodels.UserView, error) {

	_, err := GetByUsername(ctx, username)
	if err == nil {
		// found user with that username
		return nil, errUserExists
	}

	ok, coll := database.UserView(ctx)
	if !ok {
		return nil, errNotConnected
	}

	view := &dfdmodels.UserView{
		Username: username,
	}

	v, err := coll.InsertOneAndFind(ctx, view, &dfdmodels.UserView{})
	if err != nil {
		return nil, err
	}

	return v.(*dfdmodels.UserView), nil
}

// GetOrCreate creates a new UserView for the logged in user
func GetOrCreate(ctx context.Context, username string) (*dfdmodels.UserView, error) {

	view, err := GetByUsername(ctx, username)
	if err == nil {
		return view, nil
	}

	view, err = Create(ctx, username)
	if err == nil {
		return view, nil
	}

	return nil, err
}
