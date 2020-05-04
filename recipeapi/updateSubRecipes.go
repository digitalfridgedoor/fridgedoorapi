package recipeapi

import (
	"context"
	"fmt"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddSubRecipe adds a link between the recipe and the subrecipe
func (editable *EditableRecipe) AddSubRecipe(ctx context.Context, subRecipeID *primitive.ObjectID) (*Recipe, error) {

	if *editable.db.ID == *subRecipeID {
		return nil, errSubRecipe
	}

	// todo: append parent so we know when we unlink?
	if len(editable.db.ParentIds) > 0 {
		fmt.Println("Cannot add subrecipe to subrecipe")
		return nil, errSubRecipe
	}

	if editable.hasSubRecipe(subRecipeID) {
		return nil, errDuplicate
	}

	subRecipe, err := FindOneEditable(ctx, subRecipeID, editable.user)
	if err != nil {
		return nil, err
	}

	if len(subRecipe.db.Recipes) != 0 {
		return nil, errSubRecipe
	}

	subRecipe.appendParentRecipeID(*editable.db.ID)
	ok := subRecipe.save(ctx)
	if !ok {
		fmt.Println("Error updating subrecipe")
		return nil, err
	}

	editable.db.Recipes = append(editable.db.Recipes, dfdmodels.SubRecipe{
		RecipeID: *subRecipe.db.ID,
		Name:     subRecipe.db.Name,
	})

	return editable.saveAndGetDto(ctx)
}

// RemoveSubRecipe the link between the recipe/subrecipe
func (editable *EditableRecipe) RemoveSubRecipe(ctx context.Context, subRecipeID *primitive.ObjectID) (*Recipe, error) {

	filterFn := func(id *dfdmodels.SubRecipe) bool {
		return id.RecipeID != *subRecipeID
	}

	editable.filterSubRecipes(filterFn)

	subRecipe, err := FindOneEditable(ctx, subRecipeID, editable.user)

	if err == nil {
		subRecipe.removeParentRecipeID(*editable.db.ID)
		ok := subRecipe.save(ctx)
		if !ok {
			fmt.Printf("Error updating subrecipe: %v.", err)
		}
	} else {
		fmt.Printf("Could not find subrecipe with id=%v.\n", subRecipeID)
	}

	return editable.saveAndGetDto(ctx)
}

func (editable *EditableRecipe) hasSubRecipe(subRecipeID *primitive.ObjectID) bool {
	for _, subrecipe := range editable.db.Recipes {
		if subrecipe.RecipeID == *subRecipeID {
			return true
		}
	}

	return false
}

func (editable *EditableRecipe) appendParentRecipeID(parentID primitive.ObjectID) {
	hasParentID := false

	for _, id := range editable.db.ParentIds {
		if id == parentID {
			hasParentID = true
		}
	}

	if !hasParentID {
		editable.db.ParentIds = append(editable.db.ParentIds, parentID)
	}
}

func (editable *EditableRecipe) removeParentRecipeID(parentID primitive.ObjectID) {
	filtered := []primitive.ObjectID{}

	for _, id := range editable.db.ParentIds {
		if id != parentID {
			filtered = append(filtered, id)
		}
	}

	editable.db.ParentIds = filtered
}

func (editable *EditableRecipe) filterSubRecipes(filterFn func(ing *dfdmodels.SubRecipe) bool) {
	filtered := []dfdmodels.SubRecipe{}

	for _, sr := range editable.db.Recipes {
		if filterFn(&sr) {
			filtered = append(filtered, sr)
		}
	}

	editable.db.Recipes = filtered
}
