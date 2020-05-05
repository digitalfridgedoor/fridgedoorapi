package database

import (
	"context"

	"github.com/maisiesadler/gomongo"
)

// UserView returns an ICollection for the mongodb collection users
func UserView(ctx context.Context) (bool, gomongo.ICollection) {
	return gomongo.CreateCollection(ctx, "recipeapi", "userviews")
}

// Recipe returns an ICollection for the mongodb collection recipe
func Recipe(ctx context.Context) (bool, gomongo.ICollection) {
	return gomongo.CreateCollection(ctx, "recipeapi", "recipes")
}

// Ingredient returns an ICollection for the mongodb collection ingredient
func Ingredient(ctx context.Context) (bool, gomongo.ICollection) {
	return gomongo.CreateCollection(ctx, "recipeapi", "ingredients")
}

// Plan returns an ICollection for the mongodb collection plans
func Plan(ctx context.Context) (bool, gomongo.ICollection) {
	return gomongo.CreateCollection(ctx, "recipeapi", "plans")
}
