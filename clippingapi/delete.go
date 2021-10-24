package clippingapi

import (
	"context"
	"fmt"

	"fridgedoorapi/database"
	"fridgedoorapi/fridgedoorgateway"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete removes a clipping
func Delete(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, clippingID *primitive.ObjectID) error {
	ok, coll := database.Clipping(ctx)
	if !ok {
		return errNotConnected
	}

	err := coll.DeleteByID(ctx, clippingID)
	if err != nil {
		fmt.Printf("Error deleting clipping with id '%v': %v.\n", clippingID, err)
		return err
	}

	return nil
}
