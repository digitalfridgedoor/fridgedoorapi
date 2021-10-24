package clippingapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"fridgedoorapi/dfdtesting"
	"fridgedoorapi/dfdtestingapi"
)

func TestSearch(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetClippingByNamePredicate()

	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	ctx := context.TODO()
	clippingname := "testname"

	id, err := Create(ctx, user, clippingname)
	assert.Nil(t, err)

	res, err := SearchByName(context.Background(), "te", user.ViewID, 10)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, clippingname, res[0].Name)
	assert.Equal(t, *id, *res[0].ID)
}
