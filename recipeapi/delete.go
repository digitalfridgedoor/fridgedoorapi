package recipeapi

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// DeleteRecipe removes the recipe from the collection, and then removes the recipe
func DeleteRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, collectionName string, recipeID string) error {

	rID, err := primitive.ObjectIDFromHex(recipeID)
	if err != nil {
		return err
	}

	err = recipe.Delete(ctx, &rID)
	if err != nil {
		fmt.Printf("Error deleting recipe with id '%v': %v.\n", recipeID, err)
		return err
	}

	return nil
}
