package linkeduserapi

import (
	"context"
	"time"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoorapi/search"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetPublicRecipes returns the public recipes user views for all users
func GetPublicRecipes(ctx context.Context) ([]*LinkedUser, error) {
	userViews, err := getLinkedUserViews(ctx)
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
	userViews, err := getLinkedUserViews(ctx)
	if err != nil {
		return nil, err
	}

	populated := []*LinkedUser{}
	for _, uv := range userViews {
		if *uv.ID != user.ViewID {
			linkedUser, err := populatePublicUser(ctx, uv)
			if err == nil {
				populated = append(populated, linkedUser)
			}
		}
	}

	return populated, nil
}

func getLinkedUserViews(ctx context.Context) ([]*dfdmodels.UserView, error) {

	ok, coll := database.UserView(ctx)
	if !ok {
		return nil, errNotConnected
	}

	duration3s, _ := time.ParseDuration("3s")
	findctx, cancelFunc := context.WithTimeout(ctx, duration3s)
	defer cancelFunc()

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(25)

	ch, err := coll.Find(findctx, bson.D{{}}, findOptions, &dfdmodels.UserView{})
	if err != nil {
		return nil, err
	}

	results := make([]*dfdmodels.UserView, 0)
	for i := range ch {
		results = append(results, i.(*dfdmodels.UserView))
	}
	return results, nil
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

func findRecipes(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, linkedUserID primitive.ObjectID) ([]*search.RecipeDescription, error) {

	recipes, err := search.FindRecipeByTags(ctx, linkedUserID, []string{}, []string{}, 20)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func findPublicRecipes(ctx context.Context, linkedUserID primitive.ObjectID) ([]*search.RecipeDescription, error) {

	recipes, err := search.FindPublicRecipes(ctx, linkedUserID, 10)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}
