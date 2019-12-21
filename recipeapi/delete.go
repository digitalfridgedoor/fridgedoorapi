package recipeapi

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"
)

// DeleteRecipe removes the recipe from the collection, and then removes the recipe
func DeleteRecipe(ctx context.Context, request *events.APIGatewayProxyRequest, collectionName string, recipeID string) error {

	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return err
	}

	rID, err := primitive.ObjectIDFromHex(recipeID)
	if err != nil {
		return err
	}

	err = userview.RemoveRecipe(ctx, view.ID.Hex(), collectionName, rID)
	if err != nil {
		fmt.Printf("Error removing: %v.\n", err)
		return err
	}

	err = recipe.Delete(ctx, rID)
	if err != nil {
		fmt.Printf("Error deleting recipe with id '%v': %v.\n", recipeID, err)
		return err
	}

	return nil
}
