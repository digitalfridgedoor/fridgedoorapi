package search

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"github.com/stretchr/testify/assert"
)

func TestSubstring(t *testing.T) {
	sort := "-name"

	assert.Equal(t, "name", sort[1:])
}

func TestFindStartingWith(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetRecipeFindByNameOrTagsPredicate()

	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	recipeapi.CreateRecipe(context.TODO(), user, "fi_recipe")
	recipeapi.CreateRecipe(context.TODO(), user, "potatoes and chips")

	request := &FindRecipeRequest {
		StartsWith: "fi",
	}

	results, err := FindRecipe(context.Background(), user.ViewID, *request)

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(results))
}

func TestFindByTags(t *testing.T) {
	ctx := context.TODO()

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetRecipeFindByNameOrTagsPredicate()

	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	addRecipeWithTags(user, "recipe 123", []string{"hello"})

	results, err := FindRecipeByTags(ctx, user.ViewID, []string{"hello"}, []string{},  10)

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(results))
}

func TestFindByTagsAndName(t *testing.T) {
	ctx := context.TODO()

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetRecipeFindByNameOrTagsPredicate()

	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	addRecipeWithTags(user, "recipe 123", []string{"hello"})
	addRecipeWithTags(user, "recipe 456", []string{"goodbye"})
	addRecipeWithTags(user, "potatoe", []string{"goodbye"})

	request := &FindRecipeRequest {
		StartsWith: "recipe",
		Tags: []string{"goodbye"},
	}

	results, err := FindRecipe(ctx, user.ViewID, *request)

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, "recipe 456", results[0].Name)
}

func addRecipeWithTags(user *fridgedoorgateway.AuthenticatedUser, name string, tags []string) {
	ctx := context.TODO()

	recipe, _ := recipeapi.CreateRecipe(ctx, user, name)
	editable, _ := recipeapi.FindOneEditable(ctx, recipe.ID, user)

	for _, tag := range tags {
		updates := make(map[string]string)
		updates["tag_add"] = tag
		editable.UpdateMetadata(ctx, updates)
	}
}
