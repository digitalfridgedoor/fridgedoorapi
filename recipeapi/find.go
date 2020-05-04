package recipeapi

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/database"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
)

// FindOne finds a recipe by id
func FindOne(ctx context.Context, recipeID *primitive.ObjectID, user *fridgedoorgateway.AuthenticatedUser) (*ViewableRecipe, error) {

	r, err := findOneViewable(ctx, recipeID, user)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func FindOneEditable(ctx context.Context, id *primitive.ObjectID, user *fridgedoorgateway.AuthenticatedUser) (*EditableRecipe, error) {
	r, err := findOneViewable(ctx, id, user)
	if err != nil {
		return nil, err
	}

	editable := EditableRecipe(*r)

	if editable.canEdit() {
		return &editable, nil
	}

	return nil, errAuth
}

func findOneViewable(ctx context.Context, id *primitive.ObjectID, user *fridgedoorgateway.AuthenticatedUser) (*ViewableRecipe, error) {
	r, err := findOne(ctx, id)

	if err != nil {
		return nil, err
	}

	viewable := &ViewableRecipe{
		db:   r,
		user: user,
	}

	if viewable.canView() {
		return viewable, nil
	}

	return nil, errAuth
}

func findOne(ctx context.Context, id *primitive.ObjectID) (*dfdmodels.Recipe, error) {

	ok, coll := database.Recipe(ctx)
	if !ok {
		return nil, errNotConnected
	}

	r, err := coll.FindByID(ctx, id, &dfdmodels.Recipe{})

	if err != nil {
		return nil, err
	}

	re := r.(*dfdmodels.Recipe)

	return re, err
}

func mapToDto(r *dfdmodels.Recipe, user *fridgedoorgateway.AuthenticatedUser) *Recipe {
	canEdit := canEdit(r, user.ViewID)
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
