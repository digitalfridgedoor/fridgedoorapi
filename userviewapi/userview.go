package userviewapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"github.com/digitalfridgedoor/fridgedoordatabase/userview"
)

// GetUserViewByID gets a userview by id
func GetUserViewByID(ctx context.Context, userviewID string) (*View, error) {

	// todo: auth?

	view, err := userview.FindOne(ctx, userviewID)
	if err != nil {
		return nil, err
	}

	return populateUserView(ctx, view), nil
}

// GetCollectionRecipes gets the list of recipe descriptions for a collection
func GetCollectionRecipes(ctx context.Context, collection *userview.RecipeCollection) ([]*recipe.Description, error) {
	recipes, err := recipe.FindByIds(ctx, collection.Recipes)
	return recipes, err
}

// GetOtherUsersRecipes returns a collection of user views for all users
func GetOtherUsersRecipes(ctx context.Context) ([]*View, error) {
	users, err := userview.GetLinkedUserViews(ctx)
	if err != nil {
		return nil, err
	}

	populated := make([]*View, len(users))
	for idx, user := range users {
		populated[idx] = populateUserView(ctx, user)
	}

	return populated, nil
}

func populateUserView(ctx context.Context, view *userview.View) *View {
	collections := make(map[string]*RecipeCollection)
	for k, v := range view.Collections {
		recipes, err := GetCollectionRecipes(ctx, v)
		if err == nil {
			collections[k] = &RecipeCollection{
				Name:    v.Name,
				Recipes: recipes,
			}
		}
	}

	return &View{
		ID:          view.ID,
		Nickname:    view.Nickname,
		Collections: collections,
	}
}
