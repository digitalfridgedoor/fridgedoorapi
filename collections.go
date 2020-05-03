package fridgedoorapi

import (
	"context"
	"digitalfridgedoor/fridgedoordatabase/database"
)

type ingredient struct {
	c database.ICollection
}

func createIngredient(ctx context.Context) (bool, *ingredient) {
	if ok, coll := database.Ingredient(ctx); ok {
		return true, &ingredient{coll}
	}
	return false, nil
}
