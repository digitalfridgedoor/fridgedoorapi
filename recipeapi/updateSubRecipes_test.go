package recipeapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestAddSubRecipe(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()

	user := dfdtesting.CreateTestAuthenticatedUser("TestUser")

	ctx := context.Background()

	recipeName := "new recipe"
	subRecipeName := "new sub recipe"
	recipe, err := CreateRecipe(ctx, user, recipeName)
	assert.Nil(t, err)
	assert.NotNil(t, recipe)
	assert.Equal(t, recipeName, recipe.Name)

	subRecipe, err := CreateRecipe(ctx, user, subRecipeName)
	assert.Nil(t, err)
	assert.NotNil(t, subRecipe)
	assert.Equal(t, subRecipeName, subRecipe.Name)

	r, err := FindOneEditable(ctx, recipe.ID, user)
	assert.Nil(t, err)

	latestRecipe, err := r.AddSubRecipe(ctx, subRecipe.ID)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(latestRecipe.Recipes))
	latestSubRecipe := latestRecipe.Recipes[0]

	assert.Equal(t, *subRecipe.ID, latestSubRecipe.RecipeID)
	assert.Equal(t, subRecipe.Name, subRecipeName)

	// Check actual sub recipe
	latestSubRecipeMain, err := findOneViewable(ctx, subRecipe.ID, user)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(latestSubRecipeMain.db.ParentIds))
	assert.Equal(t, *r.db.ID, latestSubRecipeMain.db.ParentIds[0])

	latestRecipe, err = r.RemoveSubRecipe(ctx, subRecipe.ID)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(latestRecipe.Recipes))

	// Check actual sub recipe
	latestSubRecipeMain, err = findOneViewable(ctx, subRecipe.ID, user)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(latestSubRecipeMain.db.ParentIds))

	DeleteRecipe(ctx, user, recipe.ID)
	DeleteRecipe(ctx, user, subRecipe.ID)
}
