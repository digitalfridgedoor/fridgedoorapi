package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// FindOne finds a recipe by id
func FindOne(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, recipeID string) (*Recipe, error) {

	r, err := recipe.FindOne(ctx, recipeID)
	if err != nil {
		return nil, err
	}

	if r.CanView(user.ViewID) {
		return mapToDto(r, user), nil
	}

	return nil, errAuth
}

// FindByName finds users recipes by name
func FindByName(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, searchTerm string) ([]*Recipe, error) {

	recipes, err := recipe.FindByName(ctx, searchTerm, user.ViewID)
	if err != nil {
		return nil, err
	}

	return mapToDtos(recipes, user), nil
}

// FindByTags finds users recipes with tags
func FindByTags(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, tags []string, notTags []string) ([]*Recipe, error) {

	recipes, err := recipe.FindByTags(ctx, user.ViewID, tags, notTags)
	if err != nil {
		return nil, err
	}

	return mapToDtos(recipes, user), nil
}

func mapToDtos(r []*recipe.Recipe, user *fridgedoorgateway.AuthenticatedUser) []*Recipe {
	mapped := []*Recipe{}
	for _, v := range r {
		if v.CanView(user.ViewID) {
			mapped = append(mapped, mapToDto(v, user))
		}
	}

	return mapped
}

func mapToDto(r *recipe.Recipe, user *fridgedoorgateway.AuthenticatedUser) *Recipe {
	canEdit := r.CanEdit(user.ViewID)
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
