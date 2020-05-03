package recipeapi

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// AddIngredient adds an ingredient to a recipe
func AddIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID *primitive.ObjectID, stepIdx int, ingredientID string) (*Recipe, error) {

	ingID, err := primitive.ObjectIDFromHex(ingredientID)
	if err != nil {
		return nil, errInvalidID
	}

	ing, err := ingredient.FindOne(ctx, &ingID)
	if err != nil {
		return nil, err
	}

	recipe, err := recipe.AddIngredient(ctx, user.ViewID, recipeID, stepIdx, ingredientID, ing.Name)
	if err != nil {
		return nil, err
	}

	return mapToDto(recipe, user), nil
}

// UpdateIngredient removes an ingredient to a recipe
func UpdateIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID *primitive.ObjectID, stepIdx int, ingredientID string, updates map[string]string) (*Recipe, error) {

	recipe, err := recipe.UpdateIngredient(ctx, user.ViewID, recipeID, stepIdx, ingredientID, updates)
	if err != nil {
		return nil, err
	}

	return mapToDto(recipe, user), nil
}

// RemoveIngredient removes an ingredient to a recipe
func RemoveIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID *primitive.ObjectID, stepIdx int, ingredientID string) (*Recipe, error) {

	recipe, err := recipe.RemoveIngredient(ctx, user.ViewID, recipeID, stepIdx, ingredientID)
	if err != nil {
		return nil, err
	}

	return mapToDto(recipe, user), nil
}
