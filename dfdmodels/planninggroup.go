package dfdmodels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PlanningGroup represents a group of people who share a meal plan
type PlanningGroup struct {
	ID      *primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Name    string               `json:"name"`
	UserIDs []primitive.ObjectID `json:"userIDs"`
}
