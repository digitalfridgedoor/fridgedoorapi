package fridgedoorapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"
	"github.com/stretchr/testify/assert"
)

func TestFindForUser(t *testing.T) {
	ctx := context.Background()
	username := "TestUser"
	collectionName := "public"
	recipeName := "test-recipe"
	testUser := dfdtesting.CreateTestAuthenticatedUser(username)
	r, err := recipeapi.CreateRecipe(ctx, testUser, collectionName, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)

	view, err := userview.GetByUsername(ctx, username)
	assert.Nil(t, err)
	assert.NotNil(t, view)

	coll, ok := view.Collections[collectionName]
	assert.True(t, ok)
	assert.Equal(t, 1, len(coll.Recipes))
	recipeID := coll.Recipes[0]
	assert.Equal(t, r.ID, recipeID)

	collectionRecipes, err := userviewapi.GetCollectionRecipes(ctx, coll)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(collectionRecipes))
	collectionRecipe := collectionRecipes[0]
	assert.Equal(t, r.ID, collectionRecipe.ID)
	assert.Equal(t, recipeName, collectionRecipe.Name)

	// Cleanup
	err = recipe.Delete(ctx, r.ID)
	assert.Nil(t, err)

	dfdtesting.DeleteTestUser(testUser)
}

func TestFindByNameForUser(t *testing.T) {
	ctx := context.Background()
	username := "TestUser"
	collectionName := "public"
	recipeName := "test-recipe"
	testUser := dfdtesting.CreateTestAuthenticatedUser(username)
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

	dfdtesting.DeleteTestUser(testUser)
}
