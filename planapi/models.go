package planapi

import "go.mongodb.org/mongo-driver/bson/primitive"

// UpdateDayPlanRequest is the api request for updating a day plan
type UpdateDayPlanRequest struct {
	Month      int
	Year       int
	Day        int
	MealIndex  int
	RecipeName string
	RecipeID   *primitive.ObjectID
	MealID     *primitive.ObjectID
}

// ClearDayPlanRequest is the api request for clearing a day plan
type ClearDayPlanRequest struct {
	Month     int
	Year      int
	Day       int
	MealIndex int
}
