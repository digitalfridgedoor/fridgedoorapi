package planapi

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()

	checkExpectedDays(t, 1, 2019, 31)
	checkExpectedDays(t, 2, 2019, 28)
	checkExpectedDays(t, 3, 2019, 31)
	checkExpectedDays(t, 4, 2019, 30)
	checkExpectedDays(t, 5, 2019, 31)
	checkExpectedDays(t, 6, 2019, 30)
	checkExpectedDays(t, 7, 2019, 31)
	checkExpectedDays(t, 8, 2019, 31)
	checkExpectedDays(t, 9, 2019, 30)
	checkExpectedDays(t, 10, 2019, 31)
	checkExpectedDays(t, 11, 2019, 30)
	checkExpectedDays(t, 12, 2019, 31)
	checkExpectedDays(t, 2, 2020, 29)

	checkInvalid(t, 0, 2019)
	checkInvalid(t, 2019, 10)
	checkInvalid(t, 1, 10)
}

func TestUpdate(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetPlanFindPredicate(dfdtesting.FindPlanByMonthAndYearTestPredicate)

	user := dfdtesting.CreateTestAuthenticatedUser("TestUser")

	recipeID, _ := primitive.ObjectIDFromHex("5d8f7300a7888700270f7752")

	name := "Test Recipe"
	anotherName := "Another Name"

	request := &UpdateDayPlanRequest{
		Month:      01,
		Day:        19,
		Year:       2020,
		MealIndex:  0,
		RecipeID:   &recipeID,
		RecipeName: name,
	}

	updated, err := UpdatePlan(context.TODO(), user, request)
	assert.Nil(t, err)

	plan, err := findOne(context.TODO(), updated.ID)
	assert.Nil(t, err)
	assert.NotNil(t, plan)
	assert.Equal(t, name, plan.Days[18].Meal[0].Name)

	request.MealIndex = 1
	request.RecipeName = anotherName
	updated, err = UpdatePlan(context.TODO(), user, request)
	assert.Nil(t, err)

	plan, err = findOne(context.TODO(), updated.ID)
	assert.Nil(t, err)
	assert.NotNil(t, plan)
	assert.Equal(t, name, plan.Days[18].Meal[0].Name)
	assert.Equal(t, anotherName, plan.Days[18].Meal[1].Name)
}

func TestUpdateForGroup(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetPlanFindPredicate(dfdtesting.FindPlanByMonthAndYearForGroupTestPredicate)

	user := dfdtesting.CreateTestAuthenticatedUser("TestUser")

	planningGroupID := primitive.NewObjectID()
	recipeID := primitive.NewObjectID()

	name := "Test Recipe"
	anotherName := "Another Name"

	request := &UpdateDayPlanRequest{
		PlanningGroupID: &planningGroupID,
		Month:           01,
		Day:             19,
		Year:            2020,
		MealIndex:       0,
		RecipeID:        &recipeID,
		RecipeName:      name,
	}

	updated, err := UpdatePlan(context.TODO(), user, request)
	assert.Nil(t, err)

	plan, err := findOne(context.TODO(), updated.ID)
	assert.Nil(t, err)
	assert.NotNil(t, plan)
	assert.Equal(t, name, plan.Days[18].Meal[0].Name)

	request.MealIndex = 1
	request.RecipeName = anotherName
	updated, err = UpdatePlan(context.TODO(), user, request)
	assert.Nil(t, err)

	plan, err = findOne(context.TODO(), updated.ID)
	assert.Nil(t, err)
	assert.NotNil(t, plan)
	assert.Equal(t, name, plan.Days[18].Meal[0].Name)
	assert.Equal(t, anotherName, plan.Days[18].Meal[1].Name)
}

func TestCanRemove(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetPlanFindPredicate(dfdtesting.FindPlanByMonthAndYearTestPredicate)

	user := dfdtesting.CreateTestAuthenticatedUser("TestUser")

	recipeID, _ := primitive.ObjectIDFromHex("5d8f7300a7888700270f7752")
	clippingID, _ := primitive.ObjectIDFromHex("5d8f7300a7888700270f7753")

	name := "Test Recipe"

	request := &UpdateDayPlanRequest{
		Month:      01,
		Day:        19,
		Year:       2020,
		MealIndex:  0,
		RecipeID:   &recipeID,
		RecipeName: name,
		ClippingID: &clippingID,
	}

	updated, err := UpdatePlan(context.TODO(), user, request)
	assert.Nil(t, err)

	plan, err := findOne(context.TODO(), updated.ID)
	assert.Nil(t, err)
	assert.NotNil(t, plan)
	assert.Equal(t, clippingID, *plan.Days[18].Meal[0].ClippingID)
	assert.Equal(t, name, plan.Days[18].Meal[0].Name)
	assert.Equal(t, recipeID, *plan.Days[18].Meal[0].RecipeID)

	request = &UpdateDayPlanRequest{
		Month:     01,
		Day:       19,
		Year:      2020,
		MealIndex: 0,
	}

	updated, err = UpdatePlan(context.TODO(), user, request)
	assert.Nil(t, err)

	plan, err = findOne(context.TODO(), updated.ID)
	assert.Nil(t, err)
	assert.NotNil(t, plan)
	assert.Equal(t, (*primitive.ObjectID)(nil), plan.Days[18].Meal[0].ClippingID)
	assert.Equal(t, "", plan.Days[18].Meal[0].Name)
	assert.Equal(t, (*primitive.ObjectID)(nil), plan.Days[18].Meal[0].RecipeID)

}

func findOne(ctx context.Context, planID *primitive.ObjectID) (*dfdmodels.Plan, error) {

	ok, coll := database.Plan(ctx)
	if !ok {
		return nil, errNotConnected
	}

	plan, err := coll.FindByID(ctx, planID, &dfdmodels.Plan{})
	if err != nil {
		return nil, err
	}

	return plan.(*dfdmodels.Plan), nil
}

func checkExpectedDays(t *testing.T, month int, year int, expected int) {
	userID := primitive.NewObjectID()
	ok, p := create(userID, month, year)

	assert.True(t, ok)
	assert.NotNil(t, p)
	assert.Equal(t, expected, len(p.Days))
}

func checkInvalid(t *testing.T, month int, year int) {
	userID := primitive.NewObjectID()
	ok, p := create(userID, month, year)

	assert.False(t, ok)
	assert.Nil(t, p)
}
