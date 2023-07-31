package recipeapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"
	"github.com/digitalfridgedoor/fridgedoorapi/ingredients"

	"github.com/stretchr/testify/assert"
)

func TestUpdateIngredient(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()

	ctx := context.TODO()
	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	ingredient, err := ingredients.IngredientCollection(ctx)
	assert.Nil(t, err)

	ingname := "test ingredient"
	ing, err := ingredient.Create(ctx, ingname)
	assert.Nil(t, err)

	recipe, err := CreateRecipe(ctx, user, "recipe")
	assert.Nil(t, err)

	editable, err := FindOneEditable(ctx, recipe.ID, user)

	// Act
	r, err := editable.AddIngredient(ctx, ing.ID)
	assert.Nil(t, err)
	assert.Equal(t, "recipe", r.Name)
	assert.Equal(t, 1, len(r.Ingredients))
	assert.Equal(t, ing.ID.Hex(), r.Ingredients[0].IngredientID)
	assert.Equal(t, ingname, r.Ingredients[0].Name)
	assert.Equal(t, "", r.Ingredients[0].Amount)
	assert.Equal(t, "", r.Ingredients[0].Preperation)
	assert.Equal(t, 0, len(r.Method))

	updates := make(map[string]string)
	updates["amount"] = "amount_updated"
	updates["name"] = "name_updated"
	updates["preperation"] = "preperation_updated"

	r, err = editable.UpdateIngredient(ctx, ing.ID.Hex(), updates)
	assert.Nil(t, err)
	assert.Equal(t, "recipe", r.Name)
	assert.Equal(t, 1, len(r.Ingredients))
	assert.Equal(t, ing.ID.Hex(), r.Ingredients[0].IngredientID)
	assert.Equal(t, "name_updated", r.Ingredients[0].Name)
	assert.Equal(t, "amount_updated", r.Ingredients[0].Amount)
	assert.Equal(t, "preperation_updated", r.Ingredients[0].Preperation)
	assert.Equal(t, 0, len(r.Method))

	r, err = editable.RemoveIngredient(ctx, ing.ID.Hex())
	assert.Nil(t, err)
	assert.Equal(t, "recipe", r.Name)
	assert.Equal(t, 0, len(r.Ingredients))
	assert.Equal(t, 0, len(r.Method))
}
