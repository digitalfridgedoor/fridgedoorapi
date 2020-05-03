package search

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/assert"
)

func TestTags(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetRecipeFindByTagsPredicate()

	ctx := context.TODO()

	userID, err := primitive.ObjectIDFromHex("5d8f7300a7888700270f7752")
	recipeName := "new recipe"
	r, err := recipe.Create(ctx, userID, recipeName)

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)

	tag := primitive.NewObjectID().Hex()

	updates := make(map[string]string)
	updates["tag_add"] = tag
	_, err = recipe.UpdateMetadata(ctx, userID, r.ID, updates)
	assert.Nil(t, err)

	results, err := FindRecipeByTags(ctx, userID, []string{tag}, []string{}, 20)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(results))

	updates = make(map[string]string)
	updates["tag_remove"] = tag
	_, err = recipe.UpdateMetadata(ctx, userID, r.ID, updates)

	results, err = FindRecipeByTags(ctx, userID, []string{tag}, []string{}, 20)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(results))

	recipe.Delete(ctx, r.ID)
}

func TestNinTags(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetRecipeFindByTagsPredicate()

	ctx := context.Background()

	userID, err := primitive.ObjectIDFromHex("5d8f7300a7888700270f7752")
	recipeName := "new recipe"
	r, err := recipe.Create(ctx, userID, recipeName)
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)

	recipe.Create(ctx, userID, recipeName)

	tag := primitive.NewObjectID().Hex()

	updates := make(map[string]string)
	updates["tag_add"] = tag
	_, err = recipe.UpdateMetadata(ctx, userID, r.ID, updates)
	assert.Nil(t, err)

	results, err := recipe.FindByTags(ctx, userID, []string{}, []string{tag}, 20)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(results), 1)

	recipeInResult := false
	for _, re := range results {
		if re.ID == r.ID {
			recipeInResult = true
		}
	}

	assert.False(t, recipeInResult)

	recipe.Delete(ctx, r.ID)
}

func TestIncludeAndNinTags(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetRecipeFindByTagsPredicate()

	ctx := context.Background()

	userID, err := primitive.ObjectIDFromHex("5d8f7300a7888700270f7752")
	recipeName := "new recipe"
	r, err := recipe.Create(ctx, userID, recipeName)
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, recipeName, r.Name)

	tag := primitive.NewObjectID().Hex()
	anothertag := primitive.NewObjectID().Hex()

	updates := make(map[string]string)
	updates["tag_add"] = tag
	r, err = recipe.UpdateMetadata(ctx, userID, r.ID, updates)
	assert.Nil(t, err)
	assert.NotNil(t, r)

	results, err := FindRecipeByTags(ctx, userID, []string{tag}, []string{anothertag}, 20)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, recipeName, results[0].Name)

	updates = make(map[string]string)
	updates["tag_add"] = anothertag
	_, err = recipe.UpdateMetadata(ctx, userID, r.ID, updates)
	assert.Nil(t, err)

	results, err = FindRecipeByTags(ctx, userID, []string{tag}, []string{anothertag}, 20)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(results))

	updates = make(map[string]string)
	updates["tag_remove"] = anothertag
	_, err = recipe.UpdateMetadata(ctx, userID, r.ID, updates)
	results, err = FindRecipeByTags(ctx, userID, []string{tag}, []string{anothertag}, 20)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, recipeName, results[0].Name)

	recipe.Delete(ctx, r.ID)
}
