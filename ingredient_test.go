package fridgedoorapi

import (
	"context"
	"regexp"
	"testing"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSearchIngredient(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetIngredientFindPredicate(func(gs *dfdmodels.Ingredient, m primitive.M) bool {
		nameval := m["name"].(bson.M)
		regexval := nameval["$regex"].(primitive.Regex)

		r := regexp.MustCompile(regexval.Pattern)

		return r.MatchString(gs.Name)
	})

	ingredient, err := IngredientCollection(context.TODO())
	assert.Nil(t, err)

	ingredient.Create(context.TODO(), "test")
	ingredient.Create(context.TODO(), "one")

	ingredients, err := ingredient.FindByName(context.TODO(), "on")

	assert.Nil(t, err)
	assert.NotNil(t, ingredients)
	assert.Equal(t, 1, len(ingredients))
}

func TestFindOne(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetIngredientFindPredicate(func(gs *dfdmodels.Ingredient, m primitive.M) bool {
		nameval := m["name"].(*[]bson.M)
		regexval := (*nameval)[0]["$regex"]

		r := regexp.MustCompile(regexval.(string))

		return r.MatchString(gs.Name)
	})

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
