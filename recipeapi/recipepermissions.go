package recipeapi

import (
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CanView returns true if the given userid is permitted to view the given recipe
func CanView(r *dfdmodels.Recipe, userID primitive.ObjectID) bool {
	if r.AddedBy == userID || r.Metadata.ViewableBy.Everyone {
		return true
	}

	for _, id := range r.Metadata.ViewableBy.Users {
		if id == userID {
			return true
		}
	}

	return false
}

func canEdit(r *dfdmodels.Recipe, userID primitive.ObjectID) bool {
	return r.AddedBy == userID
}

func (r *ViewableRecipe) canView() bool {
	return CanView(r.db, r.user.ViewID)
}

func (r *EditableRecipe) canEdit() bool {
	return canEdit(r.db, r.user.ViewID)
}
