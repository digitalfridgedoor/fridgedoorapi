package fridgedoorgateway

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCanFindUserView(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindPredicate(func(uv *dfdmodels.UserView, m primitive.M) bool {
		return uv.Username == m["username"]
	})

	request := fridgedoorgatewaytesting.CreateTestAuthorizedRequest("Test")
	user, err := GetOrCreateAuthenticatedUser(context.TODO(), request)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.Equal(t, "Test", user.Username)
}
