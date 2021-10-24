package search

import (
	"context"
	"testing"

	"fridgedoorapi/dfdtesting"
	"fridgedoorapi/dfdtestingapi"
	"fridgedoorapi/recipeapi"

	"github.com/stretchr/testify/assert"
)

func TestFindStartingWith(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetRecipeFindByNamePredicate()

	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	recipeapi.CreateRecipe(context.TODO(), user, "fi_recipe")

	results, err := FindRecipeByName(context.Background(), "fi", user.ViewID, 10)

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(results))
}
