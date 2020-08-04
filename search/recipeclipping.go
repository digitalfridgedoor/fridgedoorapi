package search

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/clippingapi"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindByName finds all recipes and clippings matching startsWith
func FindByName(ctx context.Context, startsWith string, userID primitive.ObjectID, limit int64) ([]Result, error) {
	recipes, err := FindRecipeByName(ctx, startsWith, userID, limit)
	if err != nil {
		return nil, err
	}

	clippings, err := clippingapi.SearchByName(ctx, startsWith, userID, limit)
	if err != nil {
		return nil, err
	}

	rids := make(map[primitive.ObjectID]int)
	results := []Result{}

	for _, r := range recipes {
		rids[*r.ID] = 0
		results = append(results, Result{
			RecipeID: r.ID,
			Name:     r.Name,
		})
	}

	for _, c := range clippings {
		canadd := true

		if c.RecipeID != nil {
			// don't add a clipping that has been created into a recipe
			if _, ok := rids[*c.RecipeID]; ok {
				canadd = false
			}
		}

		if canadd {
			results = append(results, Result{
				ClippingID: c.ID,
				Name:       c.Name,
			})
		}
	}

	return results, nil
}
