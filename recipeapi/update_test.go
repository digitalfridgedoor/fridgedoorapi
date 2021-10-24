package recipeapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"

	"github.com/stretchr/testify/assert"
)

func TestRename(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()

	ctx := context.TODO()
	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	recipe, err := CreateRecipe(ctx, user, "recipe")
	assert.Nil(t, err)

	editable, err := FindOneEditable(ctx, recipe.ID, user)
	newRecipeName := "recipe_updated"

	// Act
	r, err := editable.Rename(ctx, newRecipeName)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, newRecipeName, r.Name)

	latest, err := findOneViewable(ctx, recipe.ID, user)
	assert.Nil(t, err)
	assert.NotNil(t, latest.db)
	assert.Equal(t, newRecipeName, latest.db.Name)
}

func TestUpdateMetadataViewableBy(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()

	ctx := context.TODO()
	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	recipe, err := CreateRecipe(ctx, user, "recipe")
	assert.Nil(t, err)
	assert.False(t, recipe.Metadata.ViewableBy.Everyone)

	editable, err := FindOneEditable(ctx, recipe.ID, user)

	// Act
	updates := make(map[string]string)
	updates["viewableby_everyone"] = "true"
	r, err := editable.UpdateMetadata(ctx, updates)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)

	latest, err := findOneViewable(ctx, recipe.ID, user)
	assert.Nil(t, err)
	assert.NotNil(t, latest.db)

	assertForBoth(&r.Metadata, &latest.db.Metadata, func(m *dfdmodels.RecipeMetadata) {
		assert.True(t, m.ViewableBy.Everyone)
	})

	// Act
	updates = make(map[string]string)
	updates["viewableby_everyone"] = "false"
	r, err = editable.UpdateMetadata(ctx, updates)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)

	latest, err = findOneViewable(ctx, recipe.ID, user)
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
	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	recipe, err := CreateRecipe(ctx, user, "recipe")
	assert.Nil(t, err)

	editable, err := FindOneEditable(ctx, recipe.ID, user)

	// Act
	updates := make(map[string]string)
	updates["link_add"] = "one"
	r, err := editable.UpdateMetadata(ctx, updates)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)

	latest, err := findOneViewable(ctx, recipe.ID, user)
	assert.Nil(t, err)
	assert.NotNil(t, latest.db)

	assertForBoth(&r.Metadata, &latest.db.Metadata, func(m *dfdmodels.RecipeMetadata) {
		assert.Len(t, m.RecipeLinks, 1)
		assert.Equal(t, "one", m.RecipeLinks[0].URL)
	})

	// Act
	updates = make(map[string]string)
	updates["link_add"] = "two"
	r, err = editable.UpdateMetadata(ctx, updates)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)

	latest, err = findOneViewable(ctx, recipe.ID, user)
	assert.Nil(t, err)
	assert.NotNil(t, latest.db)

	assertForBoth(&r.Metadata, &latest.db.Metadata, func(m *dfdmodels.RecipeMetadata) {
		assert.Len(t, m.RecipeLinks, 2)
		assert.Equal(t, "one", m.RecipeLinks[0].URL)
		assert.Equal(t, "two", m.RecipeLinks[1].URL)
	})

	// Act
	updates = make(map[string]string)
	updates["link_update_url"] = "two_updated"
	updates["link_update_linkidx"] = "1"
	r, err = editable.UpdateMetadata(ctx, updates)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)

	latest, err = findOneViewable(ctx, recipe.ID, user)
	assert.Nil(t, err)
	assert.NotNil(t, latest.db)

	assertForBoth(&r.Metadata, &latest.db.Metadata, func(m *dfdmodels.RecipeMetadata) {
		assert.Len(t, m.RecipeLinks, 2)
		assert.Equal(t, "one", m.RecipeLinks[0].URL)
		assert.Equal(t, "two_updated", m.RecipeLinks[1].URL)
	})

	// Act
	updates = make(map[string]string)
	updates["link_update_name"] = "name_for_one"
	updates["link_update_linkidx"] = "0"
	r, err = editable.UpdateMetadata(ctx, updates)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)

	latest, err = findOneViewable(ctx, recipe.ID, user)
	assert.Nil(t, err)
	assert.NotNil(t, latest.db)

	assertForBoth(&r.Metadata, &latest.db.Metadata, func(m *dfdmodels.RecipeMetadata) {
		assert.Len(t, m.RecipeLinks, 2)
		assert.Equal(t, "name_for_one", m.RecipeLinks[0].Name)
		assert.Equal(t, "one", m.RecipeLinks[0].URL)
		assert.Equal(t, "two_updated", m.RecipeLinks[1].URL)
	})
	// Act
	updates = make(map[string]string)
	updates["link_remove"] = "one"
	r, err = editable.UpdateMetadata(ctx, updates)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)

	latest, err = findOneViewable(ctx, recipe.ID, user)
	assert.Nil(t, err)
	assert.NotNil(t, latest.db)

	assertForBoth(&r.Metadata, &latest.db.Metadata, func(m *dfdmodels.RecipeMetadata) {
		assert.Len(t, m.RecipeLinks, 1)
		assert.Equal(t, "two_updated", m.RecipeLinks[0].URL)
	})
}

func assertForBoth(m1 *dfdmodels.RecipeMetadata, m2 *dfdmodels.RecipeMetadata, assertion func(metadata *dfdmodels.RecipeMetadata)) {
	assertion(m1)
	assertion(m2)
}
