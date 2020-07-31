package planapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOne finds a plan for a user by month and year
func FindOne(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, month int, year int) (*dfdmodels.Plan, error) {

	plan, _, err := getOrCreateOne(ctx, user.ViewID, month, year)
	return plan, err
}

// FindOneForGroup finds a plan for a planning group by month and year
func FindOneForGroup(ctx context.Context, planningGroupID primitive.ObjectID, month int, year int) (*dfdmodels.Plan, error) {

	plan, _, err := getOrCreateOneForGroup(ctx, planningGroupID, month, year)
	return plan, err
}

func getOrCreateOne(ctx context.Context, userID primitive.ObjectID, month int, year int) (*dfdmodels.Plan, bool, error) {

	planBson := bson.M{"month": month, "year": year, "userid": userID}
	plan, err := findPlan(ctx, planBson)
	if err != nil {
		return nil, false, err
	}

	if len(plan) == 0 {
		if ok, p := create(userID, month, year); ok {
			return p, true, nil
		}

		return nil, false, errInvalidInput
	}

	return plan[0], false, nil
}

func getOrCreateOneForGroup(ctx context.Context, planningGroupID primitive.ObjectID, month int, year int) (*dfdmodels.Plan, bool, error) {

	planBson := bson.M{"month": month, "year": year, "planninggroupid": planningGroupID}
	plan, err := findPlan(ctx, planBson)
	if err != nil {
		return nil, false, err
	}

	if len(plan) == 0 {
		if ok, p := createForGroup(planningGroupID, month, year); ok {
			return p, true, nil
		}

		return nil, false, errInvalidInput
	}

	return plan[0], false, nil
}

func findPlan(ctx context.Context, planBson bson.M) ([]*dfdmodels.Plan, error) {

	ok, coll := database.Plan(ctx)
	if !ok {
		return nil, errNotConnected
	}

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(20)

	ch, err := coll.Find(ctx, planBson, findOptions, &dfdmodels.Plan{})

	if err != nil {
		return nil, err
	}

	results := make([]*dfdmodels.Plan, 0)
	for i := range ch {
		results = append(results, i.(*dfdmodels.Plan))
	}

	return results, nil
}
