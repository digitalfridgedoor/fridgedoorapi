package ingredients

import (
	"context"
	"time"

	"fridgedoorapi/dfdmodels"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
