package planapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateWithNoRecipe(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetPlanFindPredicate(func(p *dfdmodels.Plan, m bson.M) bool {
		month := m["month"].(int)
		year := m["year"].(int)
		userid := m["userid"].(primitive.ObjectID)

		return month == p.Month && year == p.Year && userid == p.UserID
	})

	user := dfdtesting.CreateTestAuthenticatedUser("TestUser")

	name := "Test Recipe"

	request := &UpdateDayPlanRequest{
		Month:      01,
		Day:        19,
		Year:       2020,
		MealIndex:  0,
		RecipeName: name,
	}

	updated, err := UpdatePlan(context.TODO(), user, request)
	assert.Nil(t, err)

	plan, err := findOne(context.TODO(), updated.ID)
	assert.Nil(t, err)
	assert.NotNil(t, plan)
	assert.Equal(t, name, plan.Days[18].Meal[0].Name)

	createdID := plan.Days[18].Meal[0].RecipelessMealID
	assert.NotNil(t, createdID)

	meal, err := findOneRecipelessMeal(context.TODO(), createdID)

	assert.Nil(t, err)
	assert.Equal(t, name, meal.Name)
	assert.Equal(t, user.ViewID, meal.UserID)
	assert.Equal(t, 1, len(meal.PlanLink))

	_, ok := meal.PlanLink["2020_1_19"]
	assert.True(t, ok)
}

func findOneRecipelessMeal(ctx context.Context, id *primitive.ObjectID) (*dfdmodels.RecipelessMeal, error) {

	ok, coll := database.RecipelessMeal(ctx)
	if !ok {
		return nil, errNotConnected
	}

	meal, err := coll.FindByID(ctx, id, &dfdmodels.RecipelessMeal{})
	if err != nil {
		return nil, err
	}

	return meal.(*dfdmodels.RecipelessMeal), nil
}
