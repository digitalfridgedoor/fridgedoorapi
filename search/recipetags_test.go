package search

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/assert"
)

func TestTags(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetRecipeFindByTagsPredicate()

	ctx := context.TODO()

	user := fridgedoorgatewaytesting.CreateTestAuthenticatedUser("TestUser")
	recipeName := "new recipe"
	r, err := recipeapi.CreateRecipe(ctx, user, recipeName)
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)

	editable, err := recipeapi.FindOneEditable(ctx, r.ID, user)
	assert.Nil(t, err)

	tag := primitive.NewObjectID().Hex()

	updates := make(map[string]string)
	updates["tag_add"] = tag
	_, err = editable.UpdateMetadata(ctx, user, updates)
	assert.Nil(t, err)

	results, err := FindRecipeByTags(ctx, user.ViewID, []string{tag}, []string{}, 20)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(results))

	updates = make(map[string]string)
	updates["tag_remove"] = tag
	_, err = editable.UpdateMetadata(ctx, user, updates)

	results, err = FindRecipeByTags(ctx, user.ViewID, []string{tag}, []string{}, 20)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(results))

	recipeapi.DeleteRecipe(ctx, user, r.ID)
}

func TestNinTags(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetRecipeFindByTagsPredicate()

	ctx := context.Background()

	user := fridgedoorgatewaytesting.CreateTestAuthenticatedUser("TestUser")
	recipeName := "new recipe"
	r, err := recipeapi.CreateRecipe(ctx, user, recipeName)
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)

	editable, err := recipeapi.FindOneEditable(ctx, r.ID, user)
	assert.Nil(t, err)

	recipeapi.CreateRecipe(ctx, user, recipeName)

	tag := primitive.NewObjectID().Hex()

	updates := make(map[string]string)
	updates["tag_add"] = tag
	_, err = editable.UpdateMetadata(ctx, user, updates)
	assert.Nil(t, err)

	results, err := FindRecipeByTags(ctx, user.ViewID, []string{}, []string{tag}, 20)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(results), 1)

	recipeInResult := false
	for _, re := range results {
		if re.ID == r.ID {
			recipeInResult = true
		}
	}

	assert.False(t, recipeInResult)

	recipeapi.DeleteRecipe(ctx, user, r.ID)
}

func TestIncludeAndNinTags(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetRecipeFindByTagsPredicate()

	ctx := context.Background()

	user := fridgedoorgatewaytesting.CreateTestAuthenticatedUser("TestUser")
	recipeName := "new recipe"
	r, err := recipeapi.CreateRecipe(ctx, user, recipeName)
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)

	editable, err := recipeapi.FindOneEditable(ctx, r.ID, user)
	assert.Nil(t, err)

	tag := primitive.NewObjectID().Hex()
	anothertag := primitive.NewObjectID().Hex()

	updates := make(map[string]string)
	updates["tag_add"] = tag
	r, err = editable.UpdateMetadata(ctx, user, updates)
	assert.Nil(t, err)
	assert.NotNil(t, r)

	results, err := FindRecipeByTags(ctx, user.ViewID, []string{tag}, []string{anothertag}, 20)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, recipeName, results[0].Name)

	updates = make(map[string]string)
	updates["tag_add"] = anothertag
	_, err = editable.UpdateMetadata(ctx, user, updates)
	assert.Nil(t, err)

	results, err = FindRecipeByTags(ctx, user.ViewID, []string{tag}, []string{anothertag}, 20)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(results))

	updates = make(map[string]string)
	updates["tag_remove"] = anothertag
	_, err = editable.UpdateMetadata(ctx, user, updates)
	results, err = FindRecipeByTags(ctx, user.ViewID, []string{tag}, []string{anothertag}, 20)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, recipeName, results[0].Name)

	recipeapi.DeleteRecipe(ctx, user, r.ID)
}
