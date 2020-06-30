package planapi

import (
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
)

func createMealWithNoRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, name string, year int, month int, day int) (*primitive.ObjectID, error) {
	ok, coll := database.RecipelessMeal(ctx)
	if !ok {
		return nil, errNotConnected
	}

	planKey := strconv.Itoa(year) + "_" + strconv.Itoa(month) + "_" + strconv.Itoa(day)

	planLinks := make(map[string]*dfdmodels.PlanLink)
	planLinks[planKey] = &dfdmodels.PlanLink{}

	tempRecipe := &dfdmodels.RecipelessMeal{
		UserID:   user.ViewID,
		Name:     name,
		PlanLink: planLinks,
	}

	objID, err := coll.InsertOne(ctx, tempRecipe)

	return objID, err
}
