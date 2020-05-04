package fridgedoorapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/search"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"
	"github.com/stretchr/testify/assert"
)

func TestFindForUser(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetRecipeFindByTagsPredicate()

	ctx := context.Background()
	username := "TestUser"
	recipeName := "test-recipe"
	testUser := fridgedoorgatewaytesting.CreateTestAuthenticatedUser(username)
	r, err := recipeapi.CreateRecipe(ctx, testUser, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)

	view, err := userview.GetByUsername(ctx, username)
	assert.Nil(t, err)
	assert.NotNil(t, view)

	userRecipes, err := search.FindRecipeByTags(ctx, testUser.ViewID, []string{}, []string{}, 20)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(userRecipes))
	userRecipe := userRecipes[0]
	assert.Equal(t, r.ID, userRecipe.ID)
	assert.Equal(t, recipeName, userRecipe.Name)

	// Cleanup
	err = recipeapi.DeleteRecipe(ctx, testUser, r.ID)
	assert.Nil(t, err)

	fridgedoorgatewaytesting.DeleteTestUser(testUser)
}

func TestFindByNameForUser(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetRecipeFindByNamePredicate()

	ctx := context.Background()
	username := "TestUser"
	recipeName := "test-recipe"
	testUser := fridgedoorgatewaytesting.CreateTestAuthenticatedUser(username)
	r, err := recipeapi.CreateRecipe(ctx, testUser, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)

	recipes, err := search.FindRecipeByName(ctx, "test", testUser.ViewID, 20)
	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Equal(t, 1, len(recipes))

	// Cleanup
	err = recipeapi.DeleteRecipe(ctx, testUser, r.ID)
	assert.Nil(t, err)

	fridgedoorgatewaytesting.DeleteTestUser(testUser)
}
