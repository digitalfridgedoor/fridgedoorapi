package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/ingredient"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

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
