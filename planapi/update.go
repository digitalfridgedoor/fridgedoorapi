package planapi

import (
	"context"
	"fmt"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
)

// UpdatePlan adds/updates a meal plan for the day
func UpdatePlan(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, updateRequest *UpdateDayPlanRequest) (*dfdmodels.Plan, error) {

	plan, isNew, err := getOrCreateOne(ctx, user.ViewID, updateRequest.Month, updateRequest.Year)
	if err != nil {
		return nil, err
	}

	if len(plan.Days) < updateRequest.Day {
		fmt.Printf("Invalid day (%v) for month with %v days.\n", updateRequest.Day, len(plan.Days))
		return nil, errInvalidInput
	}

	currentPlanLength := len(plan.Days[updateRequest.Day-1].Meal)

	if currentPlanLength == 0 {
		plan.Days[updateRequest.Day-1].Meal = make([]dfdmodels.Meal, updateRequest.MealIndex+1)
	} else if currentPlanLength <= updateRequest.MealIndex {
		diff := updateRequest.MealIndex + 1 - currentPlanLength
		plan.Days[updateRequest.Day-1].Meal = append(plan.Days[updateRequest.Day-1].Meal, make([]dfdmodels.Meal, diff)...)
	}

	plan.Days[updateRequest.Day-1].Meal[updateRequest.MealIndex].Name = updateRequest.RecipeName
	plan.Days[updateRequest.Day-1].Meal[updateRequest.MealIndex].RecipeID = updateRequest.RecipeID

	return addOrUpdate(ctx, isNew, plan)
}

func addOrUpdate(ctx context.Context, isNew bool, plan *dfdmodels.Plan) (*dfdmodels.Plan, error) {
	ok, coll := database.Plan(ctx)
	if !ok {
		return nil, errNotConnected
	}

	if isNew {
		n, err := coll.InsertOneAndFind(ctx, plan, &dfdmodels.Plan{})
		if err != nil {
			return nil, err
		}

		return n.(*dfdmodels.Plan), nil
	}

	err := coll.UpdateByID(ctx, plan.ID, plan)
	if err != nil {
		return nil, err
	}

	n, err := coll.FindByID(ctx, plan.ID, &dfdmodels.Plan{})
	if err != nil {
		return nil, err
	}

	return n.(*dfdmodels.Plan), nil
}
