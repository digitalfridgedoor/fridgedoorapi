package fridgedoorapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestFindOne(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetIngredientFindPredicate(dfdtesting.FindIngredientByNameTestPredicate)

	ingredient, err := IngredientCollection(context.TODO())
	assert.Nil(t, err)

	name := "name"
	i, err := ingredient.Create(context.TODO(), name)
	assert.Nil(t, err)

	ing, err := ingredient.FindOne(context.Background(), i.ID)

	assert.Nil(t, err)
	assert.NotNil(t, ing)
	assert.Equal(t, i.ID, ing.ID)
	assert.Equal(t, name, ing.Name)
}
