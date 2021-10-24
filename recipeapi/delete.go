package recipeapi

import (
	"context"
	"fmt"

	"fridgedoorapi/database"
	"fridgedoorapi/fridgedoorgateway"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteRecipe removes the recipe from the collection, and then removes the recipe
func DeleteRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID *primitive.ObjectID) error {

	ok, coll := database.Recipe(ctx)
	if !ok {
		fmt.Println("Not connected")
		return errNotConnected
	}

	err := coll.DeleteByID(ctx, recipeID)
	if err != nil {
		fmt.Printf("Error deleting recipe with id '%v': %v.\n", recipeID, err)
		return err
	}

	return nil
}
