package userviewapi

import (
	"context"
	"errors"
	"fmt"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"

	"github.com/aws/aws-lambda-go/events"
)

var errNotLoggedIn = errors.New("No user logged in")

// GetOrCreateUserView creates a new UserView for the logged in user
func GetOrCreateUserView(ctx context.Context, request *events.APIGatewayProxyRequest) (*userview.View, error) {

	username, ok := fridgedoorapi.ParseUsername(request)
	if !ok {
		return nil, errNotLoggedIn
	}

	view, err := userview.GetByUsername(ctx, username)
	if err != nil {
		view, err = userview.Create(ctx, username)
	}

	nickname, ok := fridgedoorapi.ParseNickname(request)
	if ok {
		fmt.Printf("Got nickname: %v\n", nickname)
		err = userview.SetNickname(ctx, view, nickname)
		if err != nil {
			fmt.Printf("Error setting nickname: %v\n", err)
		}
	}
	return view, nil
}

// GetUserViewByID gets a userview by id
func GetUserViewByID(ctx context.Context, userviewID string) (*userview.View, error) {

	// todo: auth?

	view, err := userview.FindOne(ctx, userviewID)
	if err == nil {
		return view, nil
	}

	return view, nil
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
