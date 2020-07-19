package clippingapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create creates a recipeless meal with the given name
func Create(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, name string) (*primitive.ObjectID, error) {
	ok, coll := database.Clipping(ctx)
	if !ok {
		return nil, errNotConnected
	}

	meal := &dfdmodels.Clipping{
		UserID: user.ViewID,
		Name:   name,
	}

	objID, err := coll.InsertOne(ctx, meal)

	return objID, err
}
