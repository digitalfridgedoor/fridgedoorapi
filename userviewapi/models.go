package userviewapi

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// View represents a users set of lists
type View struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Nickname string             `json:"nickname"`
	Tags     []string           `json:"tags"`
}
