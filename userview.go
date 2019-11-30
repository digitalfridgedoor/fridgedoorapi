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

	u, err := UserView()
	if err != nil {
		return nil, err
	}

	username, ok := ParseUsername(request)
	if !ok {
		return nil, errNotLoggedIn
	}

	userview, err := u.GetByUsername(ctx, username)
	if err == nil {
		return userview, nil
	}

	return u.Create(ctx, username)
}

// GetCollectionRecipes gets the list of recipe descriptions for a collection
func GetCollectionRecipes(ctx context.Context, collection *userview.RecipeCollection) ([]*recipe.Description, error) {

	r, err := Recipe()
	if err != nil {
		return nil, err
	}

	recipes, err := r.FindByIds(ctx, collection.Recipes)
	return recipes, err
}
