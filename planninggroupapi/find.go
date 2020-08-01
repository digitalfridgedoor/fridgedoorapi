package planningroupapi

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

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
