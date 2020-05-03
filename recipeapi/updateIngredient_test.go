package recipeapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"
	"github.com/stretchr/testify/assert"
)

func TestUpdateIngredient(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()

	ctx := context.TODO()
	user := fridgedoorgatewaytesting.CreateTestAuthenticatedUser("TestUser")

	ingredient, err := fridgedoorapi.IngredientCollection(ctx)
	assert.Nil(t, err)

	ingname := "test ingredient"
	ing, err := ingredient.Create(ctx, ingname)
	assert.Nil(t, err)

	recipe, err := CreateRecipe(ctx, user, "recipe")
	assert.Nil(t, err)

	AddMethodStep(ctx, user, recipe.ID, "Test Action")

	editable, err := FindOneEditable(ctx, recipe.ID, user.ViewID)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(editable.db.Method))

	// Act
	r, err := editable.AddIngredient(ctx, user, 0, ing.ID)
	assert.Nil(t, err)
	assert.Equal(t, "recipe", r.Name)
	assert.Equal(t, 1, len(editable.db.Method))
	assert.Equal(t, 1, len(editable.db.Method[0].Ingredients))
	assert.Equal(t, ing.ID.Hex(), editable.db.Method[0].Ingredients[0].IngredientID)
	assert.Equal(t, ingname, editable.db.Method[0].Ingredients[0].Name)
	assert.Equal(t, "", editable.db.Method[0].Ingredients[0].Amount)
	assert.Equal(t, "", editable.db.Method[0].Ingredients[0].Preperation)

	updates := make(map[string]string)
	updates["amount"] = "amount_updated"
	updates["name"] = "name_updated"
	updates["preperation"] = "preperation_updated"

	r, err = editable.UpdateIngredient(ctx, user, 0, ing.ID.Hex(), updates)
	assert.Nil(t, err)
	assert.Equal(t, "recipe", r.Name)
	assert.Equal(t, 1, len(editable.db.Method))
	assert.Equal(t, 1, len(editable.db.Method[0].Ingredients))
	assert.Equal(t, ing.ID.Hex(), editable.db.Method[0].Ingredients[0].IngredientID)
	assert.Equal(t, "name_updated", editable.db.Method[0].Ingredients[0].Name)
	assert.Equal(t, "amount_updated", editable.db.Method[0].Ingredients[0].Amount)
	assert.Equal(t, "preperation_updated", editable.db.Method[0].Ingredients[0].Preperation)

	r, err = editable.RemoveIngredient(ctx, user, 0, ing.ID.Hex())
	assert.Nil(t, err)
	assert.Equal(t, "recipe", r.Name)
	assert.Equal(t, 1, len(editable.db.Method))
	assert.Equal(t, 0, len(editable.db.Method[0].Ingredients))
}
