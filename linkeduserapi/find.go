package linkeduserapi

import (
	"context"
	"digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"digitalfridgedoor/fridgedoordatabase/recipe"
	"digitalfridgedoor/fridgedoordatabase/userview"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetOtherUsersRecipes returns a collection of user views for all users
func GetOtherUsersRecipes(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser) ([]*LinkedUser, error) {
	userViews, err := userview.GetLinkedUserViews(ctx)
	if err != nil {
		return nil, err
	}

	populated := make([]*LinkedUser, len(userViews))
	for idx, uv := range userViews {
		linkedUser, err := populateLinkedUser(ctx, user, uv)
		if err != nil {
			populated[idx] = linkedUser
		}
	}

	return populated, nil
}

func populateLinkedUser(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, view *userview.View) (*LinkedUser, error) {
	recipes, err := findRecipes(ctx, user, view.ID)
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

	return mapToDtos(recipes, user), nil
}

func mapToDtos(r []*recipe.Recipe, user *fridgedoorgateway.AuthenticatedUser) []*Recipe {
	mapped := []*Recipe{}
	for idx, v := range r {
		if v.CanView(user.ViewID) {
			mapped[idx] = mapToDto(v)
		}
	}

	return mapped
}

func mapToDto(r *recipe.Recipe) *Recipe {
	return &Recipe{
		ID:    r.ID,
		Name:  r.Name,
		Image: r.Metadata.Image,
	}
}
