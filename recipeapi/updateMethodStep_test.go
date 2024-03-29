package recipeapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"

	"github.com/stretchr/testify/assert"
)

func TestUpdateMethodStep(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()

	ctx := context.TODO()
	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	recipe, err := CreateRecipe(ctx, user, "recipe")
	assert.Nil(t, err)

	editable, err := FindOneEditable(ctx, recipe.ID, user)

	// Act
	r, err := editable.AddMethodStep(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 1, len(r.Method))
	assert.Equal(t, "", r.Method[0].Action)
	assert.Equal(t, "", r.Method[0].Description)
	assert.Equal(t, "", r.Method[0].Time)

	updates := make(map[string]string)
	updates["action"] = "action_updated"
	updates["description"] = "description_updated"
	updates["time"] = "time_updated"

	r, err = editable.UpdateMethodStep(ctx, 0, updates)
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 1, len(r.Method))
	assert.Equal(t, "action_updated", r.Method[0].Action)
	assert.Equal(t, "description_updated", r.Method[0].Description)
	assert.Equal(t, "time_updated", r.Method[0].Time)

	r, err = editable.RemoveMethodStep(ctx, 0)
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 0, len(r.Method))
}
