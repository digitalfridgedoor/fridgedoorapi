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

	plan, isNew, err := getOrCreateOneForUpdateRequest(ctx, user, updateRequest)
	if err != nil {
		return nil, err
	}

	if len(plan.Days) < updateRequest.Day {
		fmt.Printf("Invalid day (%v) for month with %v days.\n", updateRequest.Day, len(plan.Days))
		return nil, errInvalidInput
	}

	dayIdx := updateRequest.Day - 1

	currentPlanLength := len(plan.Days[dayIdx].Meal)

	if currentPlanLength == 0 {
		plan.Days[dayIdx].Meal = make([]dfdmodels.Meal, updateRequest.MealIndex+1)
	} else if currentPlanLength <= updateRequest.MealIndex {
		diff := updateRequest.MealIndex + 1 - currentPlanLength
		plan.Days[dayIdx].Meal = append(plan.Days[dayIdx].Meal, make([]dfdmodels.Meal, diff)...)
	}

	plan.Days[dayIdx].Meal[updateRequest.MealIndex].Name = updateRequest.RecipeName
	plan.Days[dayIdx].Meal[updateRequest.MealIndex].RecipeID = updateRequest.RecipeID
	plan.Days[dayIdx].Meal[updateRequest.MealIndex].ClippingID = updateRequest.ClippingID

	return addOrUpdate(ctx, isNew, plan)
}

func getOrCreateOneForUpdateRequest(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, updateRequest *UpdateDayPlanRequest) (*dfdmodels.Plan, bool, error) {
	if updateRequest.PlanningGroupID != nil {
		return getOrCreateOneForGroup(ctx, *updateRequest.PlanningGroupID, updateRequest.Month, updateRequest.Year)
	}

	return getOrCreateOne(ctx, user.ViewID, updateRequest.Month, updateRequest.Year)
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
