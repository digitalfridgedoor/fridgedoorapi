package fridgedoorapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"
	"github.com/stretchr/testify/assert"
)

func TestFindForUser(t *testing.T) {
	ctx := context.Background()
	username := "TestUser"
	collectionName := "public"
	recipeName := "test-recipe"
	testUser := fridgedoorgatewaytesting.CreateTestAuthenticatedUser(username)
	r, err := recipeapi.CreateRecipe(ctx, testUser, collectionName, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)

	view, err := userview.GetByUsername(ctx, username)
	assert.Nil(t, err)
	assert.NotNil(t, view)

	userRecipes, err := recipeapi.FindByTags(ctx, testUser, []string{}, []string{})
	assert.Nil(t, err)

	assert.Equal(t, 1, len(userRecipes))
	userRecipe := userRecipes[0]
	assert.Equal(t, r.ID, userRecipe.ID)
	assert.Equal(t, recipeName, userRecipe.Name)

	// Cleanup
	err = recipe.Delete(ctx, r.ID)
	assert.Nil(t, err)

	fridgedoorgatewaytesting.DeleteTestUser(testUser)
}

func TestFindByNameForUser(t *testing.T) {
	ctx := context.Background()
	username := "TestUser"
	collectionName := "public"
	recipeName := "test-recipe"
	testUser := fridgedoorgatewaytesting.CreateTestAuthenticatedUser(username)
	r, err := recipeapi.CreateRecipe(ctx, testUser, collectionName, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)

	recipes, err := recipeapi.FindByName(ctx, testUser, "test")
	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Equal(t, 1, len(recipes))

	// Cleanup
	err = recipe.Delete(ctx, r.ID)
	assert.Nil(t, err)

	fridgedoorgatewaytesting.DeleteTestUser(testUser)
}
