package fridgedoorapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchIngredient(t *testing.T) {
	ingredients, err := SearchIngredients("on")

	assert.Nil(t, err)
	assert.NotNil(t, ingredients)
	assert.Greater(t, len(ingredients), 0)
}
