package recipeapi

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// AddSubRecipe adds a link between the recipe and the subrecipe
func AddSubRecipe(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, subRecipeID string) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	err = recipe.AddSubRecipe(ctx, view.ID, recipeID, subRecipeID)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, view, recipeID)
}

// RemoveSubRecipe the link between the recipe/subrecipe
func RemoveSubRecipe(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, subRecipeID string) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	err = recipe.RemoveSubRecipe(ctx, view.ID, recipeID, subRecipeID)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, view, recipeID)
}
