package fridgedoorapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"
	"github.com/digitalfridgedoor/fridgedoorapi/ingredients"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndAddIngredient(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()

	ctx := context.Background()
	username := "TestUser"
	recipeName := "test-recipe"
	testUser := dfdtestingapi.CreateTestAuthenticatedUser(username)

	r, err := recipeapi.CreateRecipe(ctx, testUser, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)
	assert.Equal(t, len(r.Method), 0)
	assert.Equal(t, len(r.Recipes), 0)

	ingredientcoll, err := ingredients.IngredientCollection(context.TODO())
	assert.Nil(t, err)
	ingredient, err := ingredientcoll.Create(context.TODO(), "one")
	assert.Nil(t, err)

	editable, err := recipeapi.FindOneEditable(ctx, r.ID, testUser)
	assert.Nil(t, err)

	rv, err := editable.AddMethodStep(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, rv)

	rv, err = editable.AddIngredient(ctx, 0, ingredient.ID)
	assert.Nil(t, err)
	assert.Equal(t, len(rv.Method), 1)
	method := rv.Method[0]
	assert.Equal(t, len(method.Ingredients), 1)
	ing := method.Ingredients[0]
	assert.Equal(t, "one", ing.Name)

	// Cleanup
	err = recipeapi.DeleteRecipe(ctx, testUser, rv.ID)
	assert.Nil(t, err)

	assert.Nil(t, err)
	dfdtestingapi.DeleteTestUser(ctx, testUser)
}

func TestCreateAndAddThenRemoveIngredient(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()

	ctx := context.Background()
	username := "TestUser"
	recipeName := "test-recipe"
	testUser := dfdtestingapi.CreateTestAuthenticatedUser(username)
	r, err := recipeapi.CreateRecipe(ctx, testUser, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)
	assert.Equal(t, len(r.Method), 0)
	assert.Equal(t, len(r.Recipes), 0)

	ingredientcoll, err := ingredients.IngredientCollection(context.TODO())
	assert.Nil(t, err)
	ingredient, err := ingredientcoll.Create(context.TODO(), "one")
	assert.Nil(t, err)
	anotherIngredient, err := ingredientcoll.Create(context.TODO(), "two")
	assert.Nil(t, err)

	editable, err := recipeapi.FindOneEditable(ctx, r.ID, testUser)
	assert.Nil(t, err)

	rv, err := editable.AddMethodStep(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, rv)

	rv, err = editable.AddIngredient(ctx, 0, ingredient.ID)
	assert.Nil(t, err)
	assert.Equal(t, len(rv.Method), 1)
	method := rv.Method[0]
	assert.Equal(t, len(method.Ingredients), 1)
	contains(t, method.Ingredients, []string{"one"})

	rv, err = editable.AddIngredient(ctx, 0, anotherIngredient.ID)
	assert.Nil(t, err)
	assert.Equal(t, len(rv.Method), 1)
	method = rv.Method[0]
	assert.Equal(t, 2, len(method.Ingredients))
	contains(t, method.Ingredients, []string{"one", "two"})

	rv, err = editable.RemoveIngredient(ctx, 0, anotherIngredient.ID.Hex())
	assert.Nil(t, err)
	method = rv.Method[0]
	assert.Equal(t, 1, len(method.Ingredients))
	contains(t, method.Ingredients, []string{"one"})

	// Cleanup
	err = recipeapi.DeleteRecipe(ctx, testUser, r.ID)
	assert.Nil(t, err)

	assert.Nil(t, err)
	dfdtestingapi.DeleteTestUser(ctx, testUser)
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
