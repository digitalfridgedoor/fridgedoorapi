package fridgedoorapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"

	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndAddIngredient(t *testing.T) {
	ctx := context.Background()
	username := "TestUser"
	ingredientID := "5d8f739ba7888700270f775a"
	recipeName := "test-recipe"
	testUser := fridgedoorgatewaytesting.CreateTestAuthenticatedUser(username)

	r, err := recipeapi.CreateRecipe(ctx, testUser, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)
	assert.Equal(t, len(r.Method), 0)
	assert.Equal(t, len(r.Recipes), 0)

	rv, err := recipeapi.AddMethodStep(ctx, testUser, r.ID, "Test action")
	assert.Nil(t, err)
	assert.NotNil(t, rv)

	rv, err = recipeapi.AddIngredient(ctx, testUser, r.ID, 0, ingredientID)
	assert.Nil(t, err)
	assert.Equal(t, len(rv.Method), 1)
	method := rv.Method[0]
	assert.Equal(t, len(method.Ingredients), 1)
	ing := method.Ingredients[0]
	assert.Equal(t, "Red onion", ing.Name)

	// Cleanup
	err = recipe.Delete(ctx, rv.ID)
	assert.Nil(t, err)

	assert.Nil(t, err)
	fridgedoorgatewaytesting.DeleteTestUser(testUser)
}

func TestCreateAndAddThenRemoveIngredient(t *testing.T) {
	ctx := context.Background()
	username := "TestUser"
	ingredientID := "5d8f739ba7888700270f775a"
	anotherIngredientID := "5d8f746946106c8aee8cde38"
	recipeName := "test-recipe"
	testUser := fridgedoorgatewaytesting.CreateTestAuthenticatedUser(username)
	r, err := recipeapi.CreateRecipe(ctx, testUser, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)
	assert.Equal(t, len(r.Method), 0)
	assert.Equal(t, len(r.Recipes), 0)

	rv, err := recipeapi.AddMethodStep(ctx, testUser, r.ID, "Test action")
	assert.Nil(t, err)
	assert.NotNil(t, rv)

	rv, err = recipeapi.AddIngredient(ctx, testUser, r.ID, 0, ingredientID)
	assert.Nil(t, err)
	assert.Equal(t, len(rv.Method), 1)
	method := rv.Method[0]
	assert.Equal(t, len(method.Ingredients), 1)
	contains(t, method.Ingredients, []string{"Red onion"})

	rv, err = recipeapi.AddIngredient(ctx, testUser, r.ID, 0, anotherIngredientID)
	assert.Nil(t, err)
	assert.Equal(t, len(rv.Method), 1)
	method = rv.Method[0]
	assert.Equal(t, 2, len(method.Ingredients))
	contains(t, method.Ingredients, []string{"Red onion", "Red pepper"})

	rv, err = recipeapi.RemoveIngredient(ctx, testUser, r.ID, 0, anotherIngredientID)
	assert.Nil(t, err)
	method = rv.Method[0]
	assert.Equal(t, 1, len(method.Ingredients))
	contains(t, method.Ingredients, []string{"Red onion"})

	// Cleanup
	err = recipe.Delete(ctx, r.ID)
	assert.Nil(t, err)

	assert.Nil(t, err)
	fridgedoorgatewaytesting.DeleteTestUser(testUser)
}

func contains(t *testing.T, ingredients []dfdmodels.RecipeIngredient, expected []string) {
	names := make([]string, len(ingredients))
	for _, ing := range ingredients {
		names = append(names, ing.Name)
	}

	for _, ex := range expected {
		assert.Contains(t, names, ex)
	}
}
