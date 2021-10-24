package userviewapi

import (
	"context"

	"fridgedoorapi/dfdmodels"
)

// AddTag adds a tag to users list if it isn't already there
func (editable *EditableView) AddTag(ctx context.Context, tag string) error {

	hasTag := false
	for _, t := range editable.db.Tags {
		if t == tag {
			hasTag = true
		}
	}

	if !hasTag {
		editable.db.Tags = append(editable.db.Tags, tag)
	}

	ok := editable.save(ctx)
	if !ok {
		return errNotUpdated
	}

	return nil
}

// RemoveTag removes a tag from a recipe
func (editable *EditableView) RemoveTag(ctx context.Context, tag string) error {

	editable.db.Tags = filterTags(editable.db.Tags, tag)

	ok := editable.save(ctx)
	if !ok {
		return errNotUpdated
	}

	return nil
}

// SetNickname updates the users nickname
func (editable *EditableView) SetNickname(ctx context.Context, view *dfdmodels.UserView, nickname string) error {

	if nickname == "" || view.Nickname == nickname {
		return nil
	}

	editable.db.Nickname = nickname

	ok := editable.save(ctx)
	if !ok {
		return errNotUpdated
	}

	return nil
}

func filterTags(tags []string, tagToRemove string) []string {
	filtered := []string{}

	for _, tag := range tags {
		if tag != tagToRemove {
			filtered = append(filtered, tag)
		}
	}

	return filtered
}
