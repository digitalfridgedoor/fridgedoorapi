package recipeapi

import (
	"context"
	"errors"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/ingredients"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddIngredient adds a recipe scoped ingredient 
func (editable *EditableRecipe) AddIngredient(ctx context.Context, ingredientID *primitive.ObjectID) (*Recipe, error) {

	ingredient, err := ingredients.IngredientCollection(ctx)
	if err != nil {
		return nil, err
	}

	ing, err := ingredient.FindOne(ctx, ingredientID)
	if err != nil {
		return nil, err
	}

	if editable.containsIngredient(ingredientID.Hex()) {
		return nil, errors.New("Duplicate")
	}

	recipeIng := dfdmodels.RecipeIngredient{
		Name:         ing.Name,
		IngredientID: ingredientID.Hex(),
	}

	editable.db.Ingredients = append(editable.db.Ingredients, recipeIng)

	return editable.saveAndGetDto(ctx)
}

// UpdateIngredient updates a recipe scoped ingredient 
func (editable *EditableRecipe) UpdateIngredient(ctx context.Context, ingredientID string, updates map[string]string) (*Recipe, error) {

	editable.updateIngredientByID(ingredientID, updates)

	return editable.saveAndGetDto(ctx)
}

// RemoveIngredient removes a recipe scoped ingredient 
func (editable *EditableRecipe) RemoveIngredient(ctx context.Context, ingredientID string) (*Recipe, error) {

	filterFn := func(id *dfdmodels.RecipeIngredient) bool {
		return id.IngredientID != ingredientID
	}

	editable.filterIngredients(filterFn)

	return editable.saveAndGetDto(ctx)
}

func (editable *EditableRecipe) containsIngredient(ingredientID string) bool {
	for _, ing := range editable.db.Ingredients {
		if ing.IngredientID == ingredientID {
			return true
		}
	}

	return false
}

func (editable *EditableRecipe) filterIngredients(filterFn func(ing *dfdmodels.RecipeIngredient) bool) {
	filtered := []dfdmodels.RecipeIngredient{}

	for _, ing := range editable.db.Ingredients {
		if filterFn(&ing) {
			filtered = append(filtered, ing)
		}
	}

	editable.db.Ingredients = filtered
}

func (editable *EditableRecipe) updateIngredientByID(ingredientID string, updates map[string]string) {
	updated := make([]dfdmodels.RecipeIngredient, len(editable.db.Ingredients))

	for index, ing := range editable.db.Ingredients {
		if ing.IngredientID == ingredientID {
			if update, ok := updates["name"]; ok {
				ing.Name = update
			}
			if update, ok := updates["amount"]; ok {
				ing.Amount = update
			}
			if update, ok := updates["preperation"]; ok {
				ing.Preperation = update
			}
		}
		updated[index] = ing
	}

	editable.db.Ingredients = updated
}
