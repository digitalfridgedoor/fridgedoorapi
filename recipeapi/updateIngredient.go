package recipeapi

import (
	"context"
	"errors"
	"fmt"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddIngredient adds an ingredient to a recipe
func (editable *EditableRecipe) AddIngredient(ctx context.Context, stepIdx int, ingredientID *primitive.ObjectID) (*Recipe, error) {

	editableMethodStep, err := editable.getMethodStepByIdx(ctx, stepIdx)
	if err != nil {
		fmt.Printf("Error retreiving method step, %v.\n", err)
		return nil, err
	}

	ingredient, err := fridgedoorapi.IngredientCollection(ctx)
	if err != nil {
		return nil, err
	}

	ing, err := ingredient.FindOne(ctx, ingredientID)
	if err != nil {
		return nil, err
	}

	if editableMethodStep.containsIngredient(ingredientID.Hex()) {
		return nil, errors.New("Duplicate")
	}

	recipeIng := dfdmodels.RecipeIngredient{
		Name:         ing.Name,
		IngredientID: ingredientID.Hex(),
	}

	editableMethodStep.step.Ingredients = append(editableMethodStep.step.Ingredients, recipeIng)
	editable.db.Method[stepIdx] = *editableMethodStep.step

	return editable.saveAndGetDto(ctx)
}

// UpdateIngredient removes an ingredient to a recipe
func (editable *EditableRecipe) UpdateIngredient(ctx context.Context, stepIdx int, ingredientID string, updates map[string]string) (*Recipe, error) {

	editableMethodStep, err := editable.getMethodStepByIdx(ctx, stepIdx)
	if err != nil {
		fmt.Printf("Error retreiving method step, %v.\n", err)
		return nil, err
	}

	editableMethodStep.updateIngredientByID(ingredientID, updates)
	editable.db.Method[stepIdx] = *editableMethodStep.step

	return editable.saveAndGetDto(ctx)
}

// RemoveIngredient removes an ingredient to a recipe
func (editable *EditableRecipe) RemoveIngredient(ctx context.Context, stepIdx int, ingredientID string) (*Recipe, error) {

	editableMethodStep, err := editable.getMethodStepByIdx(ctx, stepIdx)
	if err != nil {
		fmt.Printf("Error retreiving method step, %v.\n", err)
		return nil, err
	}

	filterFn := func(id *dfdmodels.RecipeIngredient) bool {
		return id.IngredientID != ingredientID
	}

	editableMethodStep.filterIngredients(filterFn)
	editable.db.Method[stepIdx] = *editableMethodStep.step

	return editable.saveAndGetDto(ctx)
}

func (editable *editableMethodStep) containsIngredient(ingredientID string) bool {
	for _, ing := range editable.step.Ingredients {
		if ing.IngredientID == ingredientID {
			return true
		}
	}

	return false
}

func (editable *editableMethodStep) filterIngredients(filterFn func(ing *dfdmodels.RecipeIngredient) bool) {
	filtered := []dfdmodels.RecipeIngredient{}

	for _, ing := range editable.step.Ingredients {
		if filterFn(&ing) {
			filtered = append(filtered, ing)
		}
	}

	editable.step.Ingredients = filtered
}

func (editable *editableMethodStep) updateIngredientByID(ingredientID string, updates map[string]string) {
	updated := make([]dfdmodels.RecipeIngredient, len(editable.step.Ingredients))

	for index, ing := range editable.step.Ingredients {
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

	editable.step.Ingredients = updated
}
