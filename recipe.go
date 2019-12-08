package fridgedoorapi

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// CreateRecipe creates a new recipe with given name
func CreateRecipe(ctx context.Context, request *events.APIGatewayProxyRequest, collectionName string, name string) (*recipe.Recipe, error) {

	userview, err := GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	r, err := Recipe()

	recipe, err := r.Create(ctx, userview.ID, name)
	if err != nil {
		return nil, err
	}

	u, err := UserView()
	if err != nil {
		return nil, err
	}
	err = u.AddRecipe(ctx, userview.ID.Hex(), collectionName, recipe.ID)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

// AddMethodStep adds a new method step to a recipe
func AddMethodStep(ctx context.Context, recipeID string, action string) (*recipe.Recipe, error) {

	r, err := Recipe()
	if err != nil {
		return nil, err
	}

	err = r.AddMethodStep(ctx, recipeID, action)
	if err != nil {
		return nil, err
	}

	return r.FindOne(ctx, recipeID)
}

// RemoveMethodStep removes a method step
func RemoveMethodStep(ctx context.Context, recipeID string, stepIdx int) (*recipe.Recipe, error) {

	r, err := Recipe()
	if err != nil {
		return nil, err
	}

	err = r.RemoveMethodStepByIndex(ctx, recipeID, stepIdx)
	if err != nil {
		return nil, err
	}

	return r.FindOne(ctx, recipeID)
}

// AddIngredient adds an ingredient to a recipe
func AddIngredient(ctx context.Context, recipeID string, stepIdx int, ingredientID string) (*recipe.Recipe, error) {

	i, err := Ingredient()
	if err != nil {
		return nil, err
	}

	ing, err := i.FindOne(ctx, ingredientID)
	if err != nil {
		return nil, err
	}

	r, err := Recipe()
	if err != nil {
		return nil, err
	}

	err = r.AddIngredient(ctx, recipeID, stepIdx, ingredientID, ing.Name)
	if err != nil {
		return nil, err
	}

	return r.FindOne(ctx, recipeID)
}

// UpdateIngredient removes an ingredient to a recipe
func UpdateIngredient(ctx context.Context, recipeID string, stepIdx int, ingredientID string, updates map[string]string) (*recipe.Recipe, error) {

	r, err := Recipe()
	if err != nil {
		return nil, err
	}

	err = r.UpdateIngredient(ctx, recipeID, stepIdx, ingredientID, updates)
	if err != nil {
		return nil, err
	}

	return r.FindOne(ctx, recipeID)
}

// RemoveIngredient removes an ingredient to a recipe
func RemoveIngredient(ctx context.Context, recipeID string, stepIdx int, ingredientID string) (*recipe.Recipe, error) {

	r, err := Recipe()
	if err != nil {
		return nil, err
	}

	err = r.RemoveIngredient(ctx, recipeID, stepIdx, ingredientID)
	if err != nil {
		return nil, err
	}

	return r.FindOne(ctx, recipeID)
}
