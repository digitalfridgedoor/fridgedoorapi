package userviewapi

import (
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// View represents a users set of lists
type View struct {
	ID       *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Nickname string              `json:"nickname"`
	Tags     []string            `json:"tags"`
}

// EditableView is a view that allows users to edit the userview
type EditableView struct {
	db *dfdmodels.UserView
}
