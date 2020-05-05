package userviewapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"
	"github.com/stretchr/testify/assert"
)

func TestUpdateTags(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()

	username := "TestUser"
	ctx := context.TODO()

	view, err := GetOrCreate(ctx, username)
	user := fridgedoorgatewaytesting.CreateTestAuthenticatedUser(username)

	assert.Nil(t, err)
	assert.Equal(t, username, view.Username)

	editable, err := GetEditableAuthenticatedUserView(ctx, user)
	assert.Nil(t, err)

	viewID := view.ID
	tag := "tag"
	err = editable.AddTag(ctx, tag)
	assert.Nil(t, err)

	view, err = GetByUsername(context.Background(), username)
	assert.Nil(t, err)
	assert.Equal(t, viewID, view.ID)
	assert.Equal(t, 1, len(view.Tags))
	assert.Equal(t, tag, view.Tags[0])

	err = editable.AddTag(context.Background(), tag)
	assert.Nil(t, err)

	view, err = GetByUsername(context.Background(), username)
	assert.Nil(t, err)
	assert.Equal(t, viewID, view.ID)
	assert.Equal(t, 1, len(view.Tags))
	assert.Equal(t, tag, view.Tags[0])

	err = editable.RemoveTag(context.Background(), tag)
	assert.Nil(t, err)

	view, err = GetByUsername(context.Background(), username)
	assert.Nil(t, err)
	assert.Equal(t, viewID, view.ID)
	assert.Equal(t, 0, len(view.Tags))

	err = delete(context.TODO(), username)
	assert.Nil(t, err)

	_, err = GetByUsername(context.Background(), username)
	assert.NotNil(t, err)
}
