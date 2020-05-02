package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddSubRecipe adds a link between the recipe and the subrecipe
func AddSubRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, subRecipeID string) (*Recipe, error) {

	rID, err := primitive.ObjectIDFromHex(recipeID)
	if err != nil {
		return nil, errInvalidID
	}

	subrID, err := primitive.ObjectIDFromHex(subRecipeID)
	if err != nil {
		return nil, errInvalidID
	}

	r, err := recipe.AddSubRecipe(ctx, user.ViewID, &rID, &subrID)
	if err != nil {
		return nil, err
	}

	return mapToDto(r, user), nil
}

// RemoveSubRecipe the link between the recipe/subrecipe
func RemoveSubRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string, subRecipeID string) (*Recipe, error) {

	rID, err := primitive.ObjectIDFromHex(recipeID)
	if err != nil {
		return nil, errInvalidID
	}

	subrID, err := primitive.ObjectIDFromHex(subRecipeID)
	if err != nil {
		return nil, errInvalidID
	}

	r, err := recipe.RemoveSubRecipe(ctx, user.ViewID, &rID, &subrID)
	if err != nil {
		return nil, err
	}

	return mapToDto(r, user), nil
}
