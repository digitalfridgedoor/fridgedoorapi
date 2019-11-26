package fridgedoorapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndAddIngredient(t *testing.T) {
	ctx := context.Background()
	userID := "5d8f7300a7888700270f7752"
	ingredientID := "5d8f739ba7888700270f775a"
	recipeName := "test-recipe"
	r, err := CreateRecipe(ctx, userID, recipeName)

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
}
