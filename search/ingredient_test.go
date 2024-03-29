package search

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/ingredients"

	"github.com/stretchr/testify/assert"
)

func TestSearchIngredient(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetIngredientFindPredicate(dfdtesting.FindIngredientByNameTestPredicate)

	ingredient, err := ingredients.IngredientCollection(context.TODO())
	assert.Nil(t, err)
	ingredient.Create(context.TODO(), "test")
	ingredient.Create(context.TODO(), "one")

	// Act
	ingredients, err := FindIngredientByName(context.TODO(), "on")

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, ingredients)
	assert.Equal(t, 1, len(ingredients))
}
