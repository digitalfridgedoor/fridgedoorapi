package recipeapi

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"
)

// CreateRecipe creates a new recipe with given name
func CreateRecipe(ctx context.Context, request *events.APIGatewayProxyRequest, collectionName string, name string) (*recipe.Recipe, error) {

	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	recipe, err := recipe.Create(ctx, view.ID, name)
	if err != nil {
		return nil, err
	}

	err = userview.AddRecipe(ctx, view.ID.Hex(), collectionName, recipe.ID)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}
