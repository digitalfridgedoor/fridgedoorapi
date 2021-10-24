package planninggroupapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create adds a new planning group with the given user group
func Create(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, name string) (*primitive.ObjectID, error) {

	ok, coll := database.PlanningGroup(ctx)

	if !ok {
		return nil, errNotConnected
	}

	group := &dfdmodels.PlanningGroup{
		Name:    name,
		UserIDs: []primitive.ObjectID{user.ViewID},
	}

	return coll.InsertOne(ctx, group)
}

// AddToGroup adds the user to the given group
func AddToGroup(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, groupID primitive.ObjectID) error {

	ok, coll := database.PlanningGroup(ctx)

	if !ok {
		return errNotConnected
	}

	obj, err := coll.FindByID(ctx, &groupID, &dfdmodels.PlanningGroup{})
	if err != nil {
		return err
	}

	group := obj.(*dfdmodels.PlanningGroup)

	group.UserIDs = append(group.UserIDs, user.ViewID)

	return coll.UpdateByID(ctx, &groupID, group)
}
