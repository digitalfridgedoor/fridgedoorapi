package planapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestFindByMonthAndYear(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetPlanFindPredicate(dfdtesting.FindPlanByMonthAndYearTestPredicate)

	user := fridgedoorgatewaytesting.CreateTestAuthenticatedUser("TestUser")

	r, err := FindOne(context.Background(), user, 1, 2020)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 1, r.Month)
	assert.Equal(t, 2020, r.Year)
}
