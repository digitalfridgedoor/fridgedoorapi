package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddSubRecipe adds a link between the recipe and the subrecipe
func AddSubRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID *primitive.ObjectID, subRecipeID *primitive.ObjectID) (*Recipe, error) {

	r, err := recipe.AddSubRecipe(ctx, user.ViewID, recipeID, subRecipeID)
	if err != nil {
		return nil, err
	}

	return mapToDto(r, user), nil
}

// RemoveSubRecipe the link between the recipe/subrecipe
func RemoveSubRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID *primitive.ObjectID, subRecipeID *primitive.ObjectID) (*Recipe, error) {

	r, err := recipe.RemoveSubRecipe(ctx, user.ViewID, recipeID, subRecipeID)
	if err != nil {
		return nil, err
	}

	return mapToDto(r, user), nil
}
