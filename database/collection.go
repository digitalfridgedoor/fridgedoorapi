package database

import (
	"context"

	"github.com/maisiesadler/gomongo"
)

// UserView returns an ICollection for the mongodb collection userviews
func UserView(ctx context.Context) (bool, gomongo.ICollection) {
	return gomongo.CreateCollection(ctx, "recipeapi", "userviews")
}

// Recipe returns an ICollection for the mongodb collection recipes
func Recipe(ctx context.Context) (bool, gomongo.ICollection) {
	return gomongo.CreateCollection(ctx, "recipeapi", "recipes")
}

// Ingredient returns an ICollection for the mongodb collection ingredients
func Ingredient(ctx context.Context) (bool, gomongo.ICollection) {
	return gomongo.CreateCollection(ctx, "recipeapi", "ingredients")
}
