package fridgedoorapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindForUser(t *testing.T) {
	ctx := context.Background()
	username := "TestUser"
	categoryName := "public"
	recipeName := "test-recipe"
	request := createTestRequest(username)
	r, err := CreateRecipe(ctx, request, categoryName, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)

	u, err := UserView()
	assert.Nil(t, err)

	view, err := u.GetByUsername(ctx, username)
	assert.Nil(t, err)
	assert.NotNil(t, view)

	coll, ok := view.Collections[categoryName]
	assert.True(t, ok)
	assert.Equal(t, 1, len(coll.Recipes))
	recipeID := coll.Recipes[0]
	assert.Equal(t, r.ID, recipeID)

	collectionRecipes, err := GetCollectionRecipes(ctx, coll)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(collectionRecipes))
	collectionRecipe := collectionRecipes[0]
	assert.Equal(t, r.ID, collectionRecipe.ID)
	assert.Equal(t, recipeName, collectionRecipe.Name)

	// Cleanup
	recipeCollection, err := Recipe()

	err = recipeCollection.Delete(ctx, r.ID)
	assert.Nil(t, err)

	u.Delete(ctx, username)
}
