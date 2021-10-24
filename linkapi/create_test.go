package clippingapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"fridgedoorapi/dfdtesting"
)

func TestCanCreate(t *testing.T) {
	ctx := context.Background()

	dfdtesting.SetTestCollectionOverride()

	user := dfdtesting.CreateTestAuthenticatedUser("TestUser")

	id, err := Create(ctx, user, "Test Meal")

	assert.Nil(t, err)
	assert.NotNil(t, id)
}
