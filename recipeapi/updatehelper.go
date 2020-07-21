package recipeapi

import (
	"context"
	"errors"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func appendString(current []string, value string) []string {
	hasValue := false

	for _, v := range current {
		if v == value {
			hasValue = true
		}
	}

	if !hasValue {
		current = append(current, value)
	}

	return current
}

func removeString(current []string, removeValue string) []string {
	filtered := []string{}

	for _, v := range current {
		if v != removeValue {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func appendID(current []primitive.ObjectID, value primitive.ObjectID) []primitive.ObjectID {
	hasValue := false

	for _, v := range current {
		if v == value {
			hasValue = true
		}
	}

	if !hasValue {
		current = append(current, value)
	}

	return current
}

func removeID(current []primitive.ObjectID, removeValue primitive.ObjectID) []primitive.ObjectID {
	filtered := []primitive.ObjectID{}

	for _, v := range current {
		if v != removeValue {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func appendLink(current []string, url string) []string {
	hasValue := false

	for _, v := range current {
		if v == url {
			hasValue = true
		}
	}

	if !hasValue {
		current = append(current, url)
	}

	return current
}

func updateLink(current []string, updateIdx string, updateValue string) []string {
	if linkIdxInt, err := strconv.Atoi(updateIdx); err == nil {
		if linkIdxInt < len(current) {
			current[linkIdxInt] = updateValue
		}
	}

	return current
}

func removeLink(current []string, removeURL string) []string {
	filtered := []string{}

	for _, v := range current {
		if v != removeURL {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
