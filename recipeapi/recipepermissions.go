package recipeapi

func (r *ViewableRecipe) canView() bool {
	if r.db.AddedBy == r.user.ViewID || r.db.Metadata.ViewableBy.Everyone {
		return true
	}

	for _, id := range r.db.Metadata.ViewableBy.Users {
		if id == r.user.ViewID {
			return true
		}
	}

	return false
}

func (r *EditableRecipe) canEdit() bool {
	return r.db.AddedBy == r.user.ViewID
}
