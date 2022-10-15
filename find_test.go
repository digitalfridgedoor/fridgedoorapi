package fridgedoorapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
	"github.com/digitalfridgedoor/fridgedoorapi/search"

	"github.com/stretchr/testify/assert"
)

func TestFindForUser(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetRecipeFindByNameOrTagsPredicate()

	ctx := context.Background()
	username := "TestUser"
	recipeName := "test-recipe"
	testUser := dfdtestingapi.CreateTestAuthenticatedUser(username)
	r, err := recipeapi.CreateRecipe(ctx, testUser, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)

	userRecipes, err := search.FindRecipeByTags(ctx, testUser.ViewID, []string{}, []string{}, 20)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(userRecipes))
	userRecipe := userRecipes[0]
	assert.Equal(t, r.ID, userRecipe.ID)
	assert.Equal(t, recipeName, userRecipe.Name)

	// Cleanup
	err = recipeapi.DeleteRecipe(ctx, testUser, r.ID)
	assert.Nil(t, err)

	dfdtestingapi.DeleteTestUser(ctx, testUser)
}

func TestFindByNameForUser(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetRecipeFindByNameOrTagsPredicate()

	ctx := context.Background()
	username := "TestUser"
	recipeName := "test-recipe"
	testUser := dfdtestingapi.CreateTestAuthenticatedUser(username)
	r, err := recipeapi.CreateRecipe(ctx, testUser, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)

	request := &search.FindRecipeRequest {
		StartsWith: "test",
	}

	recipes, err := search.FindRecipe(ctx, testUser.ViewID, *request)
	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Equal(t, 1, len(recipes))

	// Cleanup
	err = recipeapi.DeleteRecipe(ctx, testUser, r.ID)
	assert.Nil(t, err)

	dfdtestingapi.DeleteTestUser(ctx, testUser)
}
