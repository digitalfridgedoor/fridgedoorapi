package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/userview"

	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"

	"github.com/aws/aws-lambda-go/events"
)

// FindOne finds a recipe by id
func FindOne(ctx context.Context, request *events.APIGatewayProxyRequest, recipeID string) (*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return nil, err
	}

	recipe, err := recipe.FindOne(ctx, recipeID)
	if err != nil {
		return nil, err
	}

	return mapToDto(recipe, view), nil
}

// FindByName finds users recipes by name
func FindByName(ctx context.Context, request *events.APIGatewayProxyRequest, searchTerm string) ([]*Recipe, error) {
	view, err := userviewapi.GetOrCreateUserView(ctx, request)
	if err != nil {
		return make([]*Recipe, 0), err
	}

	recipes, err := recipe.FindByName(ctx, searchTerm, view.ID)
	if err != nil {
		return nil, err
	}

	return mapToDtos(recipes, view), nil
}

func mapToDtos(r []*recipe.Recipe, view *userview.View) []*Recipe {
	mapped := make([]*Recipe, len(r))
	for idx, v := range r {
		mapped[idx] = mapToDto(v, view)
	}

	return mapped
}

func mapToDto(r *recipe.Recipe, view *userview.View) *Recipe {
	canEdit := r.AddedBy == view.ID
	return &Recipe{
		ID:        r.ID,
		Name:      r.Name,
		CanEdit:   canEdit,
		Method:    r.Method,
		Recipes:   r.Recipes,
		ParentIds: r.ParentIds,
		Image:     r.Image,
	}
}
