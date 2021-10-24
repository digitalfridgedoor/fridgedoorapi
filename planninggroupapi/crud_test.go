package planninggroupapi

import (
	"context"
	"testing"

	"fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestCreateAddAndFind(t *testing.T) {

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

	err = AddToGroup(ctx, user2, *gid)

	assert.Nil(t, err)

	user1groups, err := FindAll(ctx, user1)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(user1groups))
	assert.Equal(t, groupname, user1groups[0].Name)
	assert.Equal(t, 2, len(user1groups[0].UserIDs))

	user2groups, err := FindAll(ctx, user2)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(user2groups))
	assert.Equal(t, groupname, user2groups[0].Name)
	assert.Equal(t, 2, len(user2groups[0].UserIDs))
}
