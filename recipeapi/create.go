package recipeapi

import (
	"context"
	"fmt"
	"time"

	"github.com/digitalfridgedoor/fridgedoordatabase/database"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
)

// CreateRecipe creates a new recipe with given name
func CreateRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, name string) (*Recipe, error) {

	ok, coll := database.Recipe(ctx)
	if !ok {
		return nil, errNotConnected
	}

	recipe := &dfdmodels.Recipe{
		Name:    name,
		AddedOn: time.Now(),
		AddedBy: user.ViewID,
	}

	r, err := coll.InsertOneAndFind(ctx, recipe, &dfdmodels.Recipe{})
	if err != nil {
		fmt.Printf("Error creating recipe, %v\n", err)
		return nil, err
	}

	return mapToDto(r.(*dfdmodels.Recipe), user), err
}
