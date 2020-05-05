package userviewapi

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/database"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
)

// GetAuthenticatedUserView gets a userview for the authenticated user
func GetAuthenticatedUserView(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser) (*View, error) {

	view, err := findView(ctx, user.ViewID)
	if err != nil {
		return nil, err
	}

	return populateUserView(ctx, view), nil
}

// GetEditableAuthenticatedUserView gets an editable userview for the authenticated user
func GetEditableAuthenticatedUserView(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser) (*EditableView, error) {

	view, err := findView(ctx, user.ViewID)
	if err != nil {
		return nil, err
	}

	return &EditableView{
		db: view,
	}, err
}

// GetByUsername tries to get User by username
func GetByUsername(ctx context.Context, username string) (*dfdmodels.UserView, error) {

	ok, coll := database.UserView(ctx)
	if !ok {
		return nil, errNotConnected
	}

	uv, err := coll.FindOne(ctx, bson.D{primitive.E{Key: "username", Value: username}}, &dfdmodels.UserView{})

	if err != nil {
		return nil, err
	}

	return uv.(*dfdmodels.UserView), nil
}

func findView(ctx context.Context, userID primitive.ObjectID) (*dfdmodels.UserView, error) {

	// todo: auth?
	ok, coll := database.UserView(ctx)
	if !ok {
		return nil, errNotConnected
	}

	view, err := coll.FindByID(ctx, &userID, &dfdmodels.UserView{})
	if err != nil {
		return nil, err
	}

	return view.(*dfdmodels.UserView), nil
}

func populateUserView(ctx context.Context, view *dfdmodels.UserView) *View {

	return &View{
		ID:       view.ID,
		Nickname: view.Nickname,
		Tags:     view.Tags,
	}
}
