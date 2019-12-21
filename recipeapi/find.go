package recipeapi

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// FindOne finds a recipe by id
func FindOne(ctx context.Context, recipeID string) (*recipe.Recipe, error) {
	return recipe.FindOne(ctx, recipeID)
}

// FindByName finds users recipes by name
func FindByName(ctx context.Context, request *events.APIGatewayProxyRequest, searchTerm string) ([]*recipe.Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return make([]*recipe.Recipe, 0), err
	}

	return recipe.FindByName(ctx, searchTerm, view.ID)
}
