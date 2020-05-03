package recipeapi

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *ViewableRecipe) canView(userID primitive.ObjectID) bool {
	if r.db.AddedBy == userID || r.db.Metadata.ViewableBy.Everyone {
		return true
	}

	for _, id := range r.db.Metadata.ViewableBy.Users {
		if id == userID {
			return true
		}
	}

	return false
}

func (r *EditableRecipe) canEdit(userID primitive.ObjectID) bool {
	return r.db.AddedBy == userID
}
