package planninggroupapi

import (
	"context"
	"testing"

	"fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestFindOne(t *testing.T) {

	ctx := context.Background()

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetPlanningGroupFindByUser()

	user1 := dfdtesting.CreateTestAuthenticatedUser("User1")
	user2 := dfdtesting.CreateTestAuthenticatedUser("User2")
	groupname := "Planning group"

	gid, err := Create(ctx, user1, groupname)

	assert.Nil(t, err)
	assert.NotNil(t, gid)

	groupForUser1, err := FindOne(ctx, user1, *gid)

	assert.Nil(t, err)
	assert.NotNil(t, groupForUser1)
	assert.Equal(t, groupname, groupForUser1.Name)

	groupForUser2, err := FindOne(ctx, user2, *gid)

	assert.Equal(t, errNotInGroup, err)
	assert.Nil(t, groupForUser2)
}
