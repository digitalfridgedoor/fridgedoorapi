package recipeapi

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/ingredient"
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

// AddMethodStep adds a new method step to a recipe
func AddMethodStep(ctx context.Context, recipeID string, action string) (*recipe.Recipe, error) {

	err := recipe.AddMethodStep(ctx, recipeID, action)
	if err != nil {
		return nil, err
	}

	return recipe.FindOne(ctx, recipeID)
}

// UpdateMethodStep removes a method step
func UpdateMethodStep(ctx context.Context, recipeID string, stepIdx int, updates map[string]string) (*recipe.Recipe, error) {

	err := recipe.UpdateMethodStepByIndex(ctx, recipeID, stepIdx, updates)
	if err != nil {
		return nil, err
	}

	return recipe.FindOne(ctx, recipeID)
}

// RemoveMethodStep removes a method step
func RemoveMethodStep(ctx context.Context, recipeID string, stepIdx int) (*recipe.Recipe, error) {

	err := recipe.RemoveMethodStepByIndex(ctx, recipeID, stepIdx)
	if err != nil {
		return nil, err
	}

	return recipe.FindOne(ctx, recipeID)
}

// AddIngredient adds an ingredient to a recipe
func AddIngredient(ctx context.Context, recipeID string, stepIdx int, ingredientID string) (*recipe.Recipe, error) {

	ing, err := ingredient.FindOne(ctx, ingredientID)
	if err != nil {
		return nil, err
	}

	err = recipe.AddIngredient(ctx, recipeID, stepIdx, ingredientID, ing.Name)
	if err != nil {
		return nil, err
	}

	return recipe.FindOne(ctx, recipeID)
}

// UpdateIngredient removes an ingredient to a recipe
func UpdateIngredient(ctx context.Context, recipeID string, stepIdx int, ingredientID string, updates map[string]string) (*recipe.Recipe, error) {

	err := recipe.UpdateIngredient(ctx, recipeID, stepIdx, ingredientID, updates)
	if err != nil {
		return nil, err
	}

	return recipe.FindOne(ctx, recipeID)
}

// RemoveIngredient removes an ingredient to a recipe
func RemoveIngredient(ctx context.Context, recipeID string, stepIdx int, ingredientID string) (*recipe.Recipe, error) {

	err := recipe.RemoveIngredient(ctx, recipeID, stepIdx, ingredientID)
	if err != nil {
		return nil, err
	}

	return recipe.FindOne(ctx, recipeID)
}
