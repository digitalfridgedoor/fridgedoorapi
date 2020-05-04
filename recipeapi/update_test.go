package recipeapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"
	"github.com/stretchr/testify/assert"
)

func TestRename(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()

	ctx := context.TODO()
	user := fridgedoorgatewaytesting.CreateTestAuthenticatedUser("TestUser")

	recipe, err := CreateRecipe(ctx, user, "recipe")
	assert.Nil(t, err)

	editable, err := FindOneEditable(ctx, recipe.ID, user)
	newRecipeName := "recipe_updated"

	// Act
	r, err := editable.Rename(ctx, user, newRecipeName)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, newRecipeName, r.Name)

	latest, err := FindOne(ctx, recipe.ID, user)
	assert.Nil(t, err)
	assert.NotNil(t, latest.db)
	assert.Equal(t, newRecipeName, latest.db.Name)
}

func TestUpdateMetadataViewableBy(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()

	ctx := context.TODO()
	user := fridgedoorgatewaytesting.CreateTestAuthenticatedUser("TestUser")

	recipe, err := CreateRecipe(ctx, user, "recipe")
	assert.Nil(t, err)
	assert.False(t, recipe.Metadata.ViewableBy.Everyone)

	editable, err := FindOneEditable(ctx, recipe.ID, user)

	// Act
	updates := make(map[string]string)
	updates["viewableby_everyone"] = "true"
	r, err := editable.UpdateMetadata(ctx, user, updates)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)

	latest, err := FindOne(ctx, recipe.ID, user)
	assert.Nil(t, err)
	assert.NotNil(t, latest.db)

	assertForBoth(&r.Metadata, &latest.db.Metadata, func(m *dfdmodels.RecipeMetadata) {
		assert.True(t, m.ViewableBy.Everyone)
	})

	// Act
	updates = make(map[string]string)
	updates["viewableby_everyone"] = "false"
	r, err = editable.UpdateMetadata(ctx, user, updates)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)

	latest, err = FindOne(ctx, recipe.ID, user)
	assert.Nil(t, err)
	assert.NotNil(t, latest.db)

	assertForBoth(&r.Metadata, &latest.db.Metadata, func(m *dfdmodels.RecipeMetadata) {
		assert.False(t, m.ViewableBy.Everyone)
	})
}

func TestUpdateMetadataLinks(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()

	ctx := context.TODO()
	user := fridgedoorgatewaytesting.CreateTestAuthenticatedUser("TestUser")

	recipe, err := CreateRecipe(ctx, user, "recipe")
	assert.Nil(t, err)

	editable, err := FindOneEditable(ctx, recipe.ID, user)

	// Act
	updates := make(map[string]string)
	updates["link_add"] = "one"
	r, err := editable.UpdateMetadata(ctx, user, updates)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)

	latest, err := FindOne(ctx, recipe.ID, user)
	assert.Nil(t, err)
	assert.NotNil(t, latest.db)

	assertForBoth(&r.Metadata, &latest.db.Metadata, func(m *dfdmodels.RecipeMetadata) {
		assert.Len(t, m.Links, 1)
		assert.Equal(t, "one", m.Links[0])
	})

	// Act
	updates = make(map[string]string)
	updates["link_add"] = "two"
	r, err = editable.UpdateMetadata(ctx, user, updates)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)

	latest, err = FindOne(ctx, recipe.ID, user)
	assert.Nil(t, err)
	assert.NotNil(t, latest.db)

	assertForBoth(&r.Metadata, &latest.db.Metadata, func(m *dfdmodels.RecipeMetadata) {
		assert.Len(t, m.Links, 2)
		assert.Equal(t, "one", m.Links[0])
		assert.Equal(t, "two", m.Links[1])
	})

	// Act
	updates = make(map[string]string)
	updates["link_remove"] = "one"
	r, err = editable.UpdateMetadata(ctx, user, updates)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)

	latest, err = FindOne(ctx, recipe.ID, user)
	assert.Nil(t, err)
	assert.NotNil(t, latest.db)

	assertForBoth(&r.Metadata, &latest.db.Metadata, func(m *dfdmodels.RecipeMetadata) {
		assert.Len(t, m.Links, 1)
		assert.Equal(t, "two", m.Links[0])
	})
}

func assertForBoth(m1 *dfdmodels.RecipeMetadata, m2 *dfdmodels.RecipeMetadata, assertion func(metadata *dfdmodels.RecipeMetadata)) {
	assertion(m1)
	assertion(m2)
}
