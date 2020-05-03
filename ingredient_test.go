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
		nameval := m["name"].(*[]bson.M)
		regexval := (*nameval)[0]["$regex"]

		r, err := regexp.Compile(regexval)

		return r.MatchString(gs.Name)
	})

	CreateIngredient("test")
	CreateIngredient("one")

	ingredients, err := SearchIngredients("on")

	assert.Nil(t, err)
	assert.NotNil(t, ingredients)
	assert.Greater(t, 1, len(ingredients))
}

func TestFind(t *testing.T) {

	ok, coll := createIngredient(context.TODO())
	assert.True(t, ok)

	capital, err := coll.FindByName(context.Background(), "C")
	lowercase, err := coll.FindByName(context.Background(), "c")

	assert.Nil(t, err)
	assert.Equal(t, len(capital), len(lowercase))
}

func TestFindOne(t *testing.T) {

	ok, coll := createIngredient(context.TODO())
	assert.True(t, ok)

	id, err := primitive.ObjectIDFromHex("5d8f744446106c8aee8cde37")
	assert.Nil(t, err)

	ing, err := coll.findOne(context.Background(), &id)

	assert.Nil(t, err)
	assert.NotNil(t, ing)
	assert.Equal(t, "5dac764fa0b9423b0090a898", ing.ParentID.Hex())
	assert.Equal(t, "5d8f744446106c8aee8cde37", ing.ID.Hex())
	assert.Equal(t, "Chicken thighs", ing.Name)
}
