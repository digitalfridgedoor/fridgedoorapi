package planninggroupapi

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"fridgedoorapi/database"
	"fridgedoorapi/dfdmodels"
	"fridgedoorapi/fridgedoorgateway"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindAll finds all groups for a user
func FindAll(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser) ([]*dfdmodels.PlanningGroup, error) {

	ok, coll := database.PlanningGroup(ctx)
	if !ok {
		return nil, errNotConnected
	}

	findOptions := options.Find()
	findOptions.SetLimit(20)

	userbson := bson.M{"userids": bson.M{"$in": []primitive.ObjectID{user.ViewID}}}

	ch, err := coll.Find(ctx, userbson, findOptions, &dfdmodels.PlanningGroup{})
	if err != nil {
		return []*dfdmodels.PlanningGroup{}, err
	}

	results := []*dfdmodels.PlanningGroup{}

	for i := range ch {
		r := i.(*dfdmodels.PlanningGroup)

		results = append(results, r)
	}

	return results, nil
}

// FindOne finds a planning group that a user is part of
func FindOne(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, id primitive.ObjectID) (*dfdmodels.PlanningGroup, error) {

	ok, coll := database.PlanningGroup(ctx)
	if !ok {
		return nil, errNotConnected
	}

	findOptions := options.Find()
	findOptions.SetLimit(20)

	obj, err := coll.FindByID(ctx, &id, &dfdmodels.PlanningGroup{})
	if err != nil {
		return nil, err
	}

	group := obj.(*dfdmodels.PlanningGroup)

	for _, u := range group.UserIDs {
		if u == user.ViewID {
			// user is in group
			return group, nil
		}
	}

	return nil, errNotInGroup
}
