package fridgedoorapi

import (
	"context"
	"errors"

	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"

	"github.com/aws/aws-lambda-go/events"
)

var errNotLoggedIn = errors.New("No user logged in")

// GetOrCreateUserView creates a new UserView for the logged in user
func GetOrCreateUserView(ctx context.Context, request *events.APIGatewayProxyRequest) (*userview.View, error) {

	username, ok := ParseUsername(request)
	if !ok {
		return nil, errNotLoggedIn
	}

	view, err := userview.GetByUsername(ctx, username)
	if err == nil {
		return view, nil
	}

	return userview.Create(ctx, username)
}

// GetCollectionRecipes gets the list of recipe descriptions for a collection
func GetCollectionRecipes(ctx context.Context, collection *userview.RecipeCollection) ([]*recipe.Description, error) {
	recipes, err := recipe.FindByIds(ctx, collection.Recipes)
	return recipes, err
}

// GetOtherUsersRecipes returns a collection of user views for all users
func GetOtherUsersRecipes(ctx context.Context) ([]*userview.View, error) {
	return userview.GetLinkedUserViews(ctx)
}
