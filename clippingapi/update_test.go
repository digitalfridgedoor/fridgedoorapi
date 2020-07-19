package clippingapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestClippingDoesNotExistCannotUpdate(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()

	ctx := context.TODO()

	user := dfdtesting.CreateTestAuthenticatedUser("TestUser")

	clippingID := primitive.NewObjectID()
	updates := make(map[string]string)

	updated, err := Update(ctx, user, &clippingID, updates)
	assert.NotNil(t, err)
	assert.Nil(t, updated)
}

func TestCanUpdateClippingName(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()

	ctx := context.TODO()
	user := dfdtesting.CreateTestAuthenticatedUser("TestUser")
	clippingName := "Clipping Name"
	updatedClippingName := "Updated Clipping Name"

	clippingID, err := Create(ctx, user, clippingName)
	assert.Nil(t, err)

	updates := make(map[string]string)
	updates["rename"] = updatedClippingName

	meal, err := Update(ctx, user, clippingID, updates)
	assert.Nil(t, err)

	assert.Equal(t, updatedClippingName, meal.Name)
}

func TestCanAddAndRemoveLink(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()

	ctx := context.TODO()
	user := dfdtesting.CreateTestAuthenticatedUser("TestUser")
	clippingName := "Meal Name"

	linkName := "Link name"

	clippingID, err := Create(ctx, user, clippingName)
	assert.Nil(t, err)

	clipping, err := AddLink(ctx, user, clippingID, linkName, "url")
	assert.Nil(t, err)

	assert.Equal(t, 1, len(clipping.Links))
	assert.Equal(t, linkName, clipping.Links[0].Name)
	assert.Equal(t, "url", clipping.Links[0].URL)

	clipping, err = RemoveLink(ctx, user, clippingID, 0)
	assert.Nil(t, err)

	assert.Equal(t, 0, len(clipping.Links))
}
