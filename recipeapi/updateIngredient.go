package recipeapi

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/ingredient"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// AddIngredient adds an ingredient to a recipe
func AddIngredient(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, stepIdx int, ingredientID string) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	ing, err := ingredient.FindOne(ctx, ingredientID)
	if err != nil {
		return nil, err
	}

	err = recipe.AddIngredient(ctx, view.ID, recipeID, stepIdx, ingredientID, ing.Name)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, view, recipeID)
}

// UpdateIngredient removes an ingredient to a recipe
func UpdateIngredient(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, stepIdx int, ingredientID string, updates map[string]string) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	err = recipe.UpdateIngredient(ctx, view.ID, recipeID, stepIdx, ingredientID, updates)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, view, recipeID)
}

// RemoveIngredient removes an ingredient to a recipe
func RemoveIngredient(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string, stepIdx int, ingredientID string) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	err = recipe.RemoveIngredient(ctx, view.ID, recipeID, stepIdx, ingredientID)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, view, recipeID)
}
