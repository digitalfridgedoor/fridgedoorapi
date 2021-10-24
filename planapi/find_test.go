package planapi

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestFindByMonthAndYear(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetPlanFindPredicate(dfdtesting.FindPlanByMonthAndYearTestPredicate)

	user := dfdtesting.CreateTestAuthenticatedUser("TestUser")

	r, err := FindOne(context.Background(), user, 1, 2020)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 1, r.Month)
	assert.Equal(t, 2020, r.Year)
}

func TestFindByMonthAndYearForGroup(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetPlanFindPredicate(dfdtesting.FindPlanByMonthAndYearForGroupTestPredicate)

	planningGroupID := primitive.NewObjectID()

	r, err := FindOneForGroup(context.Background(), planningGroupID, 1, 2020)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 1, r.Month)
	assert.Equal(t, 2020, r.Year)
}
