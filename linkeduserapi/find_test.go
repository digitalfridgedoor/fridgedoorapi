package linkeduserapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFindLinked(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindPredicate(func(uv *dfdmodels.UserView, m bson.M) bool {
		if username, ok := m["username"]; ok {
			return username == uv.Username
		}

		return true
	})

	ctx := context.TODO()

	username := "TestUser"
	userviewapi.Create(ctx, username)
	userviewapi.Create(ctx, "linked1")
	userviewapi.Create(ctx, "linked2")
	userviewapi.Create(ctx, "linked3")

	r, err := getLinkedUserViews(context.Background())

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.GreaterOrEqual(t, len(r), 4)

	assert.True(t, userInUserViews(r, "TestUser"))
	assert.True(t, userInUserViews(r, "linked1"))
	assert.True(t, userInUserViews(r, "linked2"))
	assert.True(t, userInUserViews(r, "linked3"))
}

func userInUserViews(userviews []*dfdmodels.UserView, username string) bool {
	for _, uv := range userviews {
		if uv.Username == username {
			return true
		}
	}

	return false
}
