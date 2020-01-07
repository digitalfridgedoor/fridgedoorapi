package dfdtesting

import (
	"context"
	"digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"digitalfridgedoor/fridgedoordatabase/userview"

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

// CreateTestAuthenticatedUser creates a user view and AuthenticatedUser for the given user
func CreateTestAuthenticatedUser(username string) *fridgedoorgateway.AuthenticatedUser {
	request := CreateTestAuthorizedRequest(username)

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), request)

	if err != nil {
		panic(err)
	}

	return user
}

// DeleteTestUser deletes a user
func DeleteTestUser(user *fridgedoorgateway.AuthenticatedUser) {
	userview.Delete(context.TODO(), user.Username)
}