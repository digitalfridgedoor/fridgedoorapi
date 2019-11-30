package fridgedoorapi

import (
	"context"
	"errors"

	"github.com/digitalfridgedoor/fridgedoordatabase/userview"

	"github.com/aws/aws-lambda-go/events"
)

var errNotLoggedIn = errors.New("No user logged in")

// GetOrCreateUserView creates a new UserView for the logged in user
func GetOrCreateUserView(ctx context.Context, request *events.APIGatewayProxyRequest) (*userview.View, error) {

	u, err := UserView()
	if err != nil {
		return nil, err
	}

	username, ok := ParseUsername(request)
	if !ok {
		return nil, errNotLoggedIn
	}

	userview, err := u.GetByUsername(ctx, username)
	if err == nil {
		return userview, nil
	}

	return u.Create(ctx, username)
}
