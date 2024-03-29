package recipeapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()

	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	recipeName := "new recipe"
	recipe, err := CreateRecipe(context.Background(), user, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, recipe)
	assert.Equal(t, "new recipe", recipe.Name)

	err = DeleteRecipe(context.Background(), user, recipe.ID)

	assert.Nil(t, err)
}
