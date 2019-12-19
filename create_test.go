package fridgedoorapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndAddIngredient(t *testing.T) {
	ctx := context.Background()
	username := "TestUser"
	ingredientID := "5d8f739ba7888700270f775a"
	collectionName := "public"
	recipeName := "test-recipe"
	request := CreateTestAuthorizedRequest(username)

	r, err := CreateRecipe(ctx, request, collectionName, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)
	assert.Equal(t, len(r.Method), 0)
	assert.Equal(t, len(r.Recipes), 0)

	recipeID := r.ID.Hex()
	r, err = AddMethodStep(ctx, recipeID, "Test action")
	assert.Nil(t, err)
	assert.NotNil(t, r)

	r, err = AddIngredient(ctx, recipeID, 0, ingredientID)
	assert.Nil(t, err)
	assert.Equal(t, len(r.Method), 1)
	method := r.Method[0]
	assert.Equal(t, len(method.Ingredients), 1)
	ing := method.Ingredients[0]
	assert.Equal(t, "Red onion", ing.Name)

	// Cleanup
	err = recipe.Delete(ctx, r.ID)
	assert.Nil(t, err)

	assert.Nil(t, err)
	userview.Delete(ctx, username)
}

func TestCreateAndAddThenRemoveIngredient(t *testing.T) {
	ctx := context.Background()
	username := "TestUser"
	ingredientID := "5d8f739ba7888700270f775a"
	anotherIngredientID := "5d8f746946106c8aee8cde38"
	collectionName := "public"
	recipeName := "test-recipe"
	request := CreateTestAuthorizedRequest(username)
	r, err := CreateRecipe(ctx, request, collectionName, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)
	assert.Equal(t, len(r.Method), 0)
	assert.Equal(t, len(r.Recipes), 0)

	recipeID := r.ID.Hex()
	r, err = AddMethodStep(ctx, recipeID, "Test action")
	assert.Nil(t, err)
	assert.NotNil(t, r)

	r, err = AddIngredient(ctx, recipeID, 0, ingredientID)
	assert.Nil(t, err)
	assert.Equal(t, len(r.Method), 1)
	method := r.Method[0]
	assert.Equal(t, len(method.Ingredients), 1)
	contains(t, method.Ingredients, []string{"Red onion"})

	r, err = AddIngredient(ctx, recipeID, 0, anotherIngredientID)
	assert.Nil(t, err)
	assert.Equal(t, len(r.Method), 1)
	method = r.Method[0]
	assert.Equal(t, 2, len(method.Ingredients))
	contains(t, method.Ingredients, []string{"Red onion", "Red pepper"})

	r, err = RemoveIngredient(ctx, recipeID, 0, anotherIngredientID)
	assert.Nil(t, err)
	method = r.Method[0]
	assert.Equal(t, 1, len(method.Ingredients))
	contains(t, method.Ingredients, []string{"Red onion"})

	// Cleanup
	err = recipe.Delete(ctx, r.ID)
	assert.Nil(t, err)

	assert.Nil(t, err)
	userview.Delete(ctx, username)
}

func contains(t *testing.T, ingredients []recipe.Ingredient, expected []string) {
	names := make([]string, len(ingredients))
	for _, ing := range ingredients {
		names = append(names, ing.Name)
	}

	for _, ex := range expected {
		assert.Contains(t, names, ex)
	}
}
