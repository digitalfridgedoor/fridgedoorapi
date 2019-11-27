package fridgedoorapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// CreateRecipe creates a new recipe with given name
func CreateRecipe(ctx context.Context, userID string, name string) (*recipe.Recipe, error) {

	u, err := User()
	if err != nil {
		return nil, err
	}

	userInfo, err := u.FindOne(ctx, userID)
	if err != nil {
		return nil, err
	}

	r, err := Recipe()

	return r.Create(ctx, userInfo.ID, name)
}

// AddIngredient adds an ingredient to a recipe
func AddIngredient(ctx context.Context, recipeID string, ingredientID string) (*recipe.Recipe, error) {

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

	err = r.AddIngredient(ctx, recipeID, ingredientID, ing.Name)
	if err != nil {
		return nil, err
	}

	return r.FindOne(ctx, recipeID)
}

// RemoveIngredient removes an ingredient to a recipe
func RemoveIngredient(ctx context.Context, recipeID string, ingredientID string) (*recipe.Recipe, error) {

	r, err := Recipe()
	if err != nil {
		return nil, err
	}

	err = r.RemoveIngredient(ctx, recipeID, ingredientID)
	if err != nil {
		return nil, err
	}

	return r.FindOne(ctx, recipeID)
}
