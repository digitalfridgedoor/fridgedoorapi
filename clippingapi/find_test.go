package clippingapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"
)

func TestFind(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()

	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	ctx := context.TODO()
	clippingname := "testname"

	id, err := Create(ctx, user, clippingname)
	assert.Nil(t, err)

	r, err := FindOne(context.Background(), user, id)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, clippingname, r.Name)
	assert.Equal(t, user.ViewID, r.UserID)
}
