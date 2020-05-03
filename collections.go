package fridgedoorapi

import (
	"context"
	"digitalfridgedoor/fridgedoordatabase/database"
)

// Ingredient is the ingredient collection
type Ingredient struct {
	c database.ICollection
}

// IngredientCollection returns a connection to mongodb ingredient collection
func IngredientCollection(ctx context.Context) (bool, *Ingredient) {
	if ok, coll := database.Ingredient(ctx); ok {
		return true, &Ingredient{coll}
	}

	return false, nil
}
