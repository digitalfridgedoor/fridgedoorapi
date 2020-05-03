package planapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoordatabase/plan"
)

// FindOne finds a recipe by id
func FindOne(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, month int, year int) (*dfdmodels.Plan, error) {

	return plan.FindByMonthAndYear(ctx, user.ViewID, month, year)
}
