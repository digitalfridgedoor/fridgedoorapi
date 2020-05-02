package userviewapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"
)

// GetUserViewByID gets a userview by id
func GetUserViewByID(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser) (*View, error) {

	// todo: auth?

	view, err := userview.FindOne(ctx, user.ViewID)
	if err != nil {
		return nil, err
	}

	return populateUserView(ctx, view), nil
}

func populateUserView(ctx context.Context, view *dfdmodels.UserView) *View {

	return &View{
		ID:       view.ID,
		Nickname: view.Nickname,
		Tags:     view.Tags,
	}
}
