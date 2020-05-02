package linkeduserapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetPublicRecipes returns the public recipes user views for all users
func GetPublicRecipes(ctx context.Context) ([]*LinkedUser, error) {
	userViews, err := userview.GetLinkedUserViews(ctx)
	if err != nil {
		return nil, err
	}

	populated := make([]*LinkedUser, len(userViews))
	for idx, uv := range userViews {
		linkedUser, err := populatePublicUser(ctx, uv)
		if err == nil {
			populated[idx] = linkedUser
		}
	}

	return populated, nil
}

// GetOtherUsersRecipes returns a collection of user views for all users
func GetOtherUsersRecipes(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser) ([]*LinkedUser, error) {
	userViews, err := userview.GetLinkedUserViews(ctx)
	if err != nil {
		return nil, err
	}

	populated := make([]*LinkedUser, len(userViews))
	for idx, uv := range userViews {
		linkedUser, err := populateLinkedUser(ctx, user, uv)
		if err == nil {
			populated[idx] = linkedUser
		}
	}

	return populated, nil
}

func populateLinkedUser(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, view *dfdmodels.UserView) (*LinkedUser, error) {
	recipes, err := findRecipes(ctx, user, *view.ID)
	if err != nil {
		return nil, err
	}

	return &LinkedUser{
		ID:       view.ID,
		Nickname: view.Nickname,
		Recipes:  recipes,
	}, nil
}

func populatePublicUser(ctx context.Context, view *dfdmodels.UserView) (*LinkedUser, error) {
	recipes, err := findPublicRecipes(ctx, *view.ID)
	if err != nil {
		return nil, err
	}

	return &LinkedUser{
		ID:       view.ID,
		Nickname: view.Nickname,
		Recipes:  recipes,
	}, nil
}

func findRecipes(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, linkedUserID primitive.ObjectID) ([]*Recipe, error) {

	recipes, err := recipe.FindByTags(ctx, linkedUserID, []string{}, []string{})
	if err != nil {
		return nil, err
	}

	return mapToAuthedDtos(recipes, user), nil
}

func findPublicRecipes(ctx context.Context, linkedUserID primitive.ObjectID) ([]*Recipe, error) {

	recipes, err := recipe.FindPublic(ctx, linkedUserID)
	if err != nil {
		return nil, err
	}

	return mapToDtos(recipes), nil
}

func mapToAuthedDtos(r []*recipe.Recipe, user *fridgedoorgateway.AuthenticatedUser) []*Recipe {
	mapped := []*Recipe{}
	for _, v := range r {
		if v.CanView(*user.ViewID) {
			mapped = append(mapped, mapToDto(v))
		}
	}

	return mapped
}

func mapToDtos(r []*recipe.Recipe) []*Recipe {
	mapped := []*Recipe{}
	for _, v := range r {
		mapped = append(mapped, mapToDto(v))
	}

	return mapped
}

func mapToDto(r *recipe.Recipe) *Recipe {
	return &Recipe{
		ID:    &r.ID,
		Name:  r.Name,
		Image: r.Metadata.Image,
	}
}
