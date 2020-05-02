package fridgedoorgateway

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"
)

var errNotLoggedIn = errors.New("No user logged in")

// GetOrCreateAuthenticatedUser creates a new UserView for the logged in user
func GetOrCreateAuthenticatedUser(ctx context.Context, request *events.APIGatewayProxyRequest) (*AuthenticatedUser, error) {

	username, ok := ParseUsername(request)
	if !ok {
		return nil, errNotLoggedIn
	}

	view, err := userview.GetByUsername(ctx, username)
	if err != nil {
		view, err = userview.Create(ctx, username)
		if err != nil {
			fmt.Printf("Error creating new user: %v", err)
			return nil, err
		}
	}

	nickname, ok := parseNickname(request)
	if ok {
		fmt.Printf("Got nickname: %v\n", nickname)
		err = userview.SetNickname(ctx, view, nickname)
		if err != nil {
			fmt.Printf("Error setting nickname: %v\n", err)
		}
	}

	user := &AuthenticatedUser{
		Username: view.Username,
		ViewID:   view.ID,
	}
	return user, nil
}

// ParseUsername attempts to parse the cognito username from the Authorizer
func ParseUsername(request *events.APIGatewayProxyRequest) (string, bool) {
	if claims, ok := request.RequestContext.Authorizer["claims"]; ok {
		c := claims.(map[string]interface{})
		username, ok := c["cognito:username"]
		return username.(string), ok
	}

	return "", false
}

// parseNickname attempts to parse the username from the Authorizer
func parseNickname(request *events.APIGatewayProxyRequest) (string, bool) {
	if claims, ok := request.RequestContext.Authorizer["claims"]; ok {
		c := claims.(map[string]interface{})
		nickname, ok := c["nickname"]
		if ok {
			fmt.Printf("Got nickname: %v.\n", nickname)
			return nickname.(string), true
		}

		fmt.Printf("Could not find nickname.\n")
		return "", false
	}

	return "", false
}
