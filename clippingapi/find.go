package clippingapi

import (
	"context"

	"fridgedoorapi/database"
	"fridgedoorapi/dfdmodels"
	"fridgedoorapi/fridgedoorgateway"

	"github.com/maisiesadler/gomongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindOne finds a clipping by id
func FindOne(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, id *primitive.ObjectID) (*dfdmodels.Clipping, error) {
	ok, coll := database.Clipping(ctx)
	if !ok {
		return nil, errNotConnected
	}

	return findClippingByID(ctx, coll, id)
}

func findClippingByID(ctx context.Context, coll gomongo.ICollection, id *primitive.ObjectID) (*dfdmodels.Clipping, error) {

	obj, err := coll.FindByID(ctx, id, &dfdmodels.Clipping{})
	if err != nil {
		return nil, err
	}

	meal := obj.(*dfdmodels.Clipping)

	return meal, nil
}
