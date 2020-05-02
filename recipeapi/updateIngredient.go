package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/ingredient"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// AddIngredient adds an ingredient to a recipe
func AddIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, stepIdx int, ingredientID string) (*Recipe, error) {

	ing, err := ingredient.FindOne(ctx, ingredientID)
	if err != nil {
		return nil, err
	}

	err = recipe.AddIngredient(ctx, *user.ViewID, recipeID, stepIdx, ingredientID, ing.Name)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, user, recipeID)
}

// UpdateIngredient removes an ingredient to a recipe
func UpdateIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, stepIdx int, ingredientID string, updates map[string]string) (*Recipe, error) {

	err := recipe.UpdateIngredient(ctx, *user.ViewID, recipeID, stepIdx, ingredientID, updates)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, user, recipeID)
}

// RemoveIngredient removes an ingredient to a recipe
func RemoveIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, stepIdx int, ingredientID string) (*Recipe, error) {

	err := recipe.RemoveIngredient(ctx, *user.ViewID, recipeID, stepIdx, ingredientID)
	if err != nil {
		return nil, err
	}

	return findOneAndMap(ctx, user, recipeID)
}
