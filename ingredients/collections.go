package ingredients

import (
	"context"

	"fridgedoorapi/database"

	"github.com/maisiesadler/gomongo"
)

// Ingredient is the ingredient collection
type Ingredient struct {
	c gomongo.ICollection
}

// IngredientCollection returns a connection to mongodb ingredient collection
func IngredientCollection(ctx context.Context) (*Ingredient, error) {
	if ok, coll := database.Ingredient(ctx); ok {
		return &Ingredient{coll}, nil
	}

	return nil, errNotConnected
}
