package recipeapi

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// FindOne finds a recipe by id
func FindOne(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID *primitive.ObjectID) (*Recipe, error) {

	r, err := recipe.FindOne(ctx, recipeID, user.ViewID)
	if err != nil {
		return nil, err
	}

	if recipe.CanView(r, user.ViewID) {
		return mapToDto(r, user), nil
	}

	return nil, errAuth
}

// FindByName finds users recipes by name
func FindByName(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, searchTerm string) ([]*recipe.Description, error) {

	recipes, err := recipe.FindByName(ctx, searchTerm, user.ViewID, 20)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

// FindByTags finds users recipes with tags
func FindByTags(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, tags []string, notTags []string) ([]*recipe.Description, error) {

	recipes, err := recipe.FindByTags(ctx, user.ViewID, tags, notTags, 20)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func mapToDto(r *dfdmodels.Recipe, user *fridgedoorgateway.AuthenticatedUser) *Recipe {
	canEdit := recipe.CanEdit(r, user.ViewID)
	ownedByUser := r.AddedBy == user.ViewID
	return &Recipe{
		ID:          r.ID,
		Name:        r.Name,
		CanEdit:     canEdit,
		OwnedByUser: ownedByUser,
		Method:      r.Method,
		Recipes:     r.Recipes,
		ParentIds:   r.ParentIds,
		Metadata:    r.Metadata,
	}
}
