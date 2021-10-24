package search

import (
	"context"
	"testing"

	"fridgedoorapi/clippingapi"

	"fridgedoorapi/dfdtesting"
	"fridgedoorapi/recipeapi"

	"github.com/stretchr/testify/assert"
)

func TestFindRecipeOrClippingStartingWith(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetRecipeFindByNamePredicate()
	dfdtesting.SetClippingByNamePredicate()

	user := dfdtesting.CreateTestAuthenticatedUser("TestUser")

	cid, err := clippingapi.Create(context.TODO(), user, "fi_clipping")
	assert.Nil(t, err)
	rec, err := recipeapi.CreateRecipe(context.TODO(), user, "fi_recipe")
	assert.Nil(t, err)
	clippingapi.Create(context.TODO(), user, "another_fi_clipping")
	recipeapi.CreateRecipe(context.TODO(), user, "another_fi_recipe")

	results, err := FindByName(context.Background(), "fi", user.ViewID, 10)

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 2, len(results))
	assert.Equal(t, *rec.ID, *results[0].RecipeID)
	assert.Equal(t, "fi_recipe", results[0].Name)
	assert.Equal(t, *cid, *results[1].ClippingID)
	assert.Equal(t, "fi_clipping", results[1].Name)
}
