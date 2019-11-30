package fridgedoorapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndAddIngredient(t *testing.T) {
	ctx := context.Background()
	username := "TestUser"
	ingredientID := "5d8f739ba7888700270f775a"
	categoryName := "public"
	recipeName := "test-recipe"
	request := createTestRequest(username)
	r, err := CreateRecipe(ctx, request, categoryName, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)
	assert.Equal(t, len(r.Ingredients), 0)
	assert.Equal(t, len(r.Recipes), 0)

	r, err = AddIngredient(ctx, r.ID.Hex(), ingredientID)
	assert.Nil(t, err)
	assert.Equal(t, len(r.Ingredients), 1)
	ing := r.Ingredients[0]
	assert.Equal(t, "Red onion", ing.Name)

	// Cleanup
	recipeCollection, err := Recipe()

	err = recipeCollection.Delete(ctx, r.ID)
	assert.Nil(t, err)

	u, err := UserView()
	assert.Nil(t, err)
	u.Delete(ctx, username)
}

func TestCreateAndAddThenRemoveIngredient(t *testing.T) {
	ctx := context.Background()
	username := "TestUser"
	ingredientID := "5d8f739ba7888700270f775a"
	anotherIngredientID := "5d8f746946106c8aee8cde38"
	categoryName := "public"
	recipeName := "test-recipe"
	request := createTestRequest(username)
	r, err := CreateRecipe(ctx, request, categoryName, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)
	assert.Equal(t, len(r.Ingredients), 0)
	assert.Equal(t, len(r.Recipes), 0)

	r, err = AddIngredient(ctx, r.ID.Hex(), ingredientID)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(r.Ingredients))
	contains(t, r.Ingredients, []string{"Red onion"})

	r, err = AddIngredient(ctx, r.ID.Hex(), anotherIngredientID)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(r.Ingredients))
	contains(t, r.Ingredients, []string{"Red onion", "Red pepper"})

	r, err = RemoveIngredient(ctx, r.ID.Hex(), anotherIngredientID)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(r.Ingredients))
	contains(t, r.Ingredients, []string{"Red onion"})

	// Cleanup
	recipeCollection, err := Recipe()

	err = recipeCollection.Delete(ctx, r.ID)
	assert.Nil(t, err)

	u, err := UserView()
	assert.Nil(t, err)
	u.Delete(ctx, username)
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
