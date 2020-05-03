package fridgedoorapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/database"
)

// Ingredient is the ingredient collection
type Ingredient struct {
	c database.ICollection
}

// IngredientCollection returns a connection to mongodb ingredient collection
func IngredientCollection(ctx context.Context) (*Ingredient, error) {
	if ok, coll := database.Ingredient(ctx); ok {
		return &Ingredient{coll}, nil
	}

	return nil, errNotConnected
}
