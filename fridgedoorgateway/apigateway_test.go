package fridgedoorgateway

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
)

func TestCanFindUserView(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()

	ctx := context.TODO()
	username := "Test"
	request := createTestAuthorizedRequest(username)

	user, err := GetOrCreateAuthenticatedUser(ctx, request)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, username, user.Username)
}

func createTestAuthorizedRequest(username string) *events.APIGatewayProxyRequest {
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
