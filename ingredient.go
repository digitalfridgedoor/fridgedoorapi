package fridgedoorapi

import (
	"context"
	"time"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create creates a new ingredient with given name
func (coll *Ingredient) Create(ctx context.Context, name string) (*dfdmodels.Ingredient, error) {

	ingredient := &dfdmodels.Ingredient{
		Name:    name,
		AddedOn: time.Now(),
	}

	insertedID, err := coll.c.InsertOne(ctx, ingredient)
	if err != nil {
		return nil, err
	}

	return coll.FindOne(ctx, insertedID)
}

// FindOne finds one ingredient matching the provided id
func (coll *Ingredient) FindOne(ctx context.Context, id *primitive.ObjectID) (*dfdmodels.Ingredient, error) {

	ing, err := coll.c.FindByID(ctx, id, &dfdmodels.Ingredient{})
	if err != nil {
		return nil, err
	}

	return ing.(*dfdmodels.Ingredient), err
}

// FindByName finds ingredients starting with the given letter
func (coll *Ingredient) FindByName(ctx context.Context, startsWith string) ([]*dfdmodels.Ingredient, error) {

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
