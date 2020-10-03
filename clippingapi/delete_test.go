package clippingapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
)

func TestCanDelete(t *testing.T) {
	ctx := context.Background()

	dfdtesting.SetTestCollectionOverride()

	user := dfdtesting.CreateTestAuthenticatedUser("TestUser")

	id, err := Create(ctx, user, "Test Meal")
	assert.Nil(t, err)

	err = Delete(ctx, user, id)

	assert.Nil(t, err)
}
