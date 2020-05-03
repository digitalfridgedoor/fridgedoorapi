package recipeapi

import (
	"context"
	"fmt"

	"github.com/digitalfridgedoor/fridgedoordatabase/database"
)

func (editable *EditableRecipe) save(ctx context.Context) bool {

	ok, coll := database.Recipe(ctx)
	if !ok {
		return false
	}

	err := coll.UpdateByID(ctx, editable.db.ID, editable.db)
	if err != nil {
		fmt.Printf("Error saving recipe: %v\n", err)
		return false
	}

	return true
}

func (editable *EditableRecipe) saveAndGetDto(ctx context.Context) (*Recipe, error) {
	ok := editable.save(ctx)
	if !ok {
		fmt.Printf("Did not save update, %v.\n", errNotConnected)
		return nil, errNotConnected
	}

	updated, err := findOne(ctx, editable.db.ID)
	if err != nil {
		fmt.Printf("Error reloading recipe, %v.\n", err)
		return nil, err
	}

	return mapToDto(updated, editable.user), nil
}
