package clippingapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"
)

func TestCanCreate(t *testing.T) {
	ctx := context.Background()

	dfdtesting.SetTestCollectionOverride()

	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	id, err := Create(ctx, user, "Test Meal")

	assert.Nil(t, err)
	assert.NotNil(t, id)
}
