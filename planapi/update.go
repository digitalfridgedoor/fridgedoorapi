package planapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/plan"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
)

// UpdatePlan adds/updates a meal plan for the day
func UpdatePlan(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, updateRequest *UpdateDayPlanRequest) (*plan.Plan, error) {

	request := &plan.UpdateDayPlanRequest{
		UserID:     *user.ViewID,
		Year:       updateRequest.Year,
		Month:      updateRequest.Month,
		Day:        updateRequest.Day,
		MealIndex:  updateRequest.MealIndex,
		RecipeName: updateRequest.RecipeName,
		RecipeID:   updateRequest.RecipeID,
	}

	planID, err := plan.Update(ctx, request)
	if err != nil {
		return nil, err
	}

	return plan.FindOne(ctx, *planID)
}
