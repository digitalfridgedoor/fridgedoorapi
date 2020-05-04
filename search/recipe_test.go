package search

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestFindStartingWith(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetRecipeFindByNamePredicate()

	user := fridgedoorgatewaytesting.CreateTestAuthenticatedUser("TestUser")

	recipeapi.CreateRecipe(context.TODO(), user, "fi_recipe")

	results, err := FindRecipeByName(context.Background(), "fi", user.ViewID, 10)

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(results))
}
