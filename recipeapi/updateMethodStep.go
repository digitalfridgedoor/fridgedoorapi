package recipeapi

import (
	"context"
	"errors"
	"fmt"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
)

// AddMethodStep adds a new empty method step to a recipe
func (editable *EditableRecipe) AddMethodStep(ctx context.Context) (*Recipe, error) {
	methodStep := dfdmodels.MethodStep{}

	editable.db.Method = append(editable.db.Method, methodStep)

	return editable.saveAndGetDto(ctx)
}

// UpdateMethodStep removes a method step
func (editable *EditableRecipe) UpdateMethodStep(ctx context.Context, stepIdx int, updates map[string]string) (*Recipe, error) {

	editableMethodStep, err := editable.getMethodStepByIdx(ctx, stepIdx)
	if err != nil {
		fmt.Printf("Error retreiving method step, %v.\n", err)
		return nil, err
	}

	editableMethodStep.updateMethodStep(updates)
	editable.db.Method[stepIdx] = *editableMethodStep.step

	return editable.saveAndGetDto(ctx)
}

// RemoveMethodStep removes a method step
func (editable *EditableRecipe) RemoveMethodStep(ctx context.Context, stepIdx int) (*Recipe, error) {

	if stepIdx < 0 {
		return nil, errors.New("Invalid index")
	}

	if len(editable.db.Method) <= stepIdx {
		return nil, errors.New("Invalid index")
	}

	copy(editable.db.Method[stepIdx:], editable.db.Method[stepIdx+1:])  // Shift a[i+1:] left one index.
	editable.db.Method = editable.db.Method[:len(editable.db.Method)-1] // Truncate slice.

	return editable.saveAndGetDto(ctx)
}

func (editable *editableMethodStep) updateMethodStep(updates map[string]string) {

	if update, ok := updates["action"]; ok {
		editable.step.Action = update
	}
	if update, ok := updates["description"]; ok {
		editable.step.Description = update
	}
	if update, ok := updates["time"]; ok {
		editable.step.Time = update
	}
}
