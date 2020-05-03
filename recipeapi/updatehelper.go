package recipeapi

import (
	"context"
	"errors"
	"strconv"
)

func (editable *EditableRecipe) getMethodStepByIdx(ctx context.Context, stepIdx int) (*editableMethodStep, error) {

	if stepIdx < 0 {
		return nil, errors.New("Invalid index, " + strconv.Itoa(stepIdx))
	}

	if len(editable.db.Method) <= stepIdx {
		return nil, errors.New("Invalid index, " + strconv.Itoa(stepIdx))
	}

	methodStep := editable.db.Method[stepIdx]

	return &editableMethodStep{&methodStep}, nil
}
