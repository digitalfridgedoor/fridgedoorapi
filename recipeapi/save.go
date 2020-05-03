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
