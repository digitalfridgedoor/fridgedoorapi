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

func TestFindStartingWith(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetRecipeFindByNamePredicate()

	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	recipeapi.CreateRecipe(context.TODO(), user, "fi_recipe")
	recipeapi.CreateRecipe(context.TODO(), user, "potatoes and chips")

	results, err := FindRecipeByName(context.Background(), "fi", user.ViewID, 10)

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(results))
}

func TestFindByTags(t *testing.T) {
	ctx := context.TODO()

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetRecipeFindByTagsPredicate()

	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")

	addRecipeWithTags(user, "recipe 123", []string{"hello"})

	results, err := FindRecipeByTags(ctx, user.ViewID, []string{"hello"}, []string{},  10)

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(results))
}

func addRecipeWithTags(user *fridgedoorgateway.AuthenticatedUser, name string, tags []string) {
	ctx := context.TODO()

	recipe, _ := recipeapi.CreateRecipe(ctx, user, "fi_recipe")
	editable, _ := recipeapi.FindOneEditable(ctx, recipe.ID, user)

	for _, tag := range tags {
		updates := make(map[string]string)
		updates["tag_add"] = tag
		editable.UpdateMetadata(ctx, updates)
	}
}
