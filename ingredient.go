package fridgedoorapi

import (
	"context"
	"time"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateIngredient creates a new ingredient
func CreateIngredient(name string) (*dfdmodels.Ingredient, error) {
	ok, ingredient := createIngredient(context.TODO())
	if !ok {
		return nil, errNotConnected
	}
	return ingredient.create(context.Background(), name)
}

// SearchIngredients retrieves the ingredients matching the query
func SearchIngredients(startsWith string) ([]*dfdmodels.Ingredient, error) {
	ok, ingredient := createIngredient(context.TODO())
	if !ok {
		return nil, errNotConnected
	}
	return ingredient.findByName(context.Background(), startsWith)
}

// Create creates a new ingredient with given name
func (coll *ingredient) create(ctx context.Context, name string) (*dfdmodels.Ingredient, error) {

	ingredient := &dfdmodels.Ingredient{
		Name:    name,
		AddedOn: time.Now(),
	}

	insertedID, err := coll.c.InsertOne(ctx, ingredient)
	if err != nil {
		return nil, err
	}

	return coll.findOne(ctx, insertedID)
}

// FindByName finds ingredients starting with the given letter
func (coll *ingredient) findByName(ctx context.Context, startsWith string) ([]*dfdmodels.Ingredient, error) {

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(20)

	regex := bson.M{"$regex": primitive.Regex{Pattern: "\\b" + startsWith, Options: "i"}}
	startsWithBson := bson.M{"name": regex}

	ch, err := coll.c.Find(ctx, startsWithBson, findOptions, &dfdmodels.Ingredient{})
	if err != nil {
		return make([]*dfdmodels.Ingredient, 0), err
	}

	results := make([]*dfdmodels.Ingredient, 0)

	for i := range ch {
		results = append(results, i.(*dfdmodels.Ingredient))
	}

	return results, nil
}

func (coll *ingredient) findOne(ctx context.Context, id *primitive.ObjectID) (*dfdmodels.Ingredient, error) {

	ing, err := coll.c.FindByID(ctx, id, &dfdmodels.Ingredient{})

	return ing.(*dfdmodels.Ingredient), err
}
