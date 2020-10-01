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
	updates["name"] = updatedClippingName

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

func TestCanUpdateLink(t *testing.T) {

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

	updatedLinkName := "updatedname"

	clipping, err = UpdateLink(ctx, user, clippingID, 0, "name", updatedLinkName)
	assert.Nil(t, err)

	assert.Equal(t, updatedLinkName, clipping.Links[0].Name)

	updatedUrl := "updatedurl"

	clipping, err = UpdateLink(ctx, user, clippingID, 0, "url", updatedUrl)
	assert.Nil(t, err)

	assert.Equal(t, updatedUrl, clipping.Links[0].URL)
}

func TestReorderLinks(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()

	ctx := context.TODO()
	user := dfdtesting.CreateTestAuthenticatedUser("TestUser")
	clippingName := "Meal Name"

	linkName1 := "Link name 1"
	linkName2 := "Link name 2"

	clippingID, err := Create(ctx, user, clippingName)
	assert.Nil(t, err)

	clipping, err := AddLink(ctx, user, clippingID, linkName1, "url1")
	assert.Nil(t, err)
	clipping, err = AddLink(ctx, user, clippingID, linkName2, "url2")
	assert.Nil(t, err)

	assert.Equal(t, 2, len(clipping.Links))
	assert.Equal(t, linkName1, clipping.Links[0].Name)
	assert.Equal(t, linkName2, clipping.Links[1].Name)

	clipping, err = SwapLinkPosition(ctx, user, clippingID, 0, 1)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(clipping.Links))
	assert.Equal(t, linkName2, clipping.Links[0].Name)
	assert.Equal(t, linkName1, clipping.Links[1].Name)
}
