package recipeapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Rename updates the name of the recipe
func (editable *EditableRecipe) Rename(ctx context.Context, name string) (*Recipe, error) {

	editable.db.Name = name

	return editable.saveAndGetDto(ctx)
}

// UpdateMetadata updates the recipes metadata property
func (editable *EditableRecipe) UpdateMetadata(ctx context.Context, updates map[string]string) (*Recipe, error) {

	editable.updateRecipeMetadata(ctx, updates)

	latest, err := editable.saveAndGetDto(ctx)
	if err != nil {
		return nil, err
	}

	if update, ok := updates["tag_add"]; ok {
		editableuser, err := userviewapi.GetEditableByID(ctx, editable.user.ViewID)
		if err == nil {
			editableuser.AddTag(ctx, update)
		}
	}

	return latest, nil
}

func (editable *EditableRecipe) updateRecipeMetadata(ctx context.Context, updates map[string]string) {

	if update, ok := updates["image"]; ok {
		editable.db.Metadata.Image = update == "true"
	}
	if update, ok := updates["tag_add"]; ok {
		editable.db.Metadata.Tags = appendString(editable.db.Metadata.Tags, update)
	}
	if update, ok := updates["tag_remove"]; ok {
		editable.db.Metadata.Tags = removeString(editable.db.Metadata.Tags, update)
	}
	if update, ok := updates["link_add"]; ok {
		editable.db.Metadata.Links = appendLink(editable.db.Metadata.Links, update)
	}
	if update, ok := updates["link_update"]; ok {
		if linkIdx, linkIdxOk := updates["link_update_linkidx"]; linkIdxOk {
			editable.db.Metadata.Links = updateLink(editable.db.Metadata.Links, linkIdx, update)
		}
	}
	if update, ok := updates["link_remove"]; ok {
		editable.db.Metadata.Links = removeLink(editable.db.Metadata.Links, update)
	}
	if update, ok := updates["viewableby_everyone"]; ok {
		editable.db.Metadata.ViewableBy.Everyone = update == "true"
	}
	if update, ok := updates["viewableby_adduser"]; ok {
		objectID, err := primitive.ObjectIDFromHex(update)
		if err == nil {
			editable.db.Metadata.ViewableBy.Users = appendID(editable.db.Metadata.ViewableBy.Users, objectID)
		}
	}
	if update, ok := updates["viewableby_removeuser"]; ok {
		objectID, err := primitive.ObjectIDFromHex(update)
		if err == nil {
			editable.db.Metadata.ViewableBy.Users = removeID(editable.db.Metadata.ViewableBy.Users, objectID)
		}
	}
}
