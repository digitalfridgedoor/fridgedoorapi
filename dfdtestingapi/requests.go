package dfdtestingapi

import (
	"context"

	"fridgedoorapi/database"
	"fridgedoorapi/fridgedoorgateway"
	"fridgedoorapi/userviewapi"

	"github.com/aws/aws-lambda-go/events"
)

// CreateTestAuthorizedRequest creates an authenticated api gateway request for the given user
func CreateTestAuthorizedRequest(username string) *events.APIGatewayProxyRequest {
	claims := make(map[string]interface{})
	claims["cognito:username"] = username
	authorizer := make(map[string]interface{})
	authorizer["claims"] = claims
	context := events.APIGatewayProxyRequestContext{
		Authorizer: authorizer,
	}
	request := &events.APIGatewayProxyRequest{
		RequestContext: context,
	}

	return request
}

// CreateTestAuthenticatedUserAndRequest creates a request, user view and AuthenticatedUser for the given user
func CreateTestAuthenticatedUserAndRequest(username string) (*fridgedoorgateway.AuthenticatedUser, *events.APIGatewayProxyRequest) {
	request := CreateTestAuthorizedRequest(username)

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), request)

	if err != nil {
		panic(err)
	}

	return user, request
}

// CreateTestAuthenticatedUser creates a user view and AuthenticatedUser for the given user
func CreateTestAuthenticatedUser(username string) *fridgedoorgateway.AuthenticatedUser {
	request := CreateTestAuthorizedRequest(username)

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), request)

	if err != nil {
		panic(err)
	}

	return user
}

// DeleteUserForRequest deletes a userview for a test request
func DeleteUserForRequest(ctx context.Context, request *events.APIGatewayProxyRequest) {
	username, ok := fridgedoorgateway.ParseUsername(request)

	if ok {
		deleteByUsername(ctx, username)
	}
}

// DeleteTestUser deletes a user
func DeleteTestUser(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser) {
	deleteByUsername(ctx, user.Username)
}

func deleteByUsername(ctx context.Context, username string) {
	ok, coll := database.UserView(ctx)
	if !ok {
		return
	}

	view, err := userviewapi.GetByUsername(ctx, username)
	if err != nil {
		return
	}

	coll.DeleteByID(ctx, view.ID)
}
