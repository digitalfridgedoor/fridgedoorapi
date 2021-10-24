package search

import (
	"context"

	"fridgedoorapi/database"
	"fridgedoorapi/dfdmodels"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindIngredientByName finds ingredients starting with the given string
func FindIngredientByName(ctx context.Context, startsWith string) ([]*dfdmodels.Ingredient, error) {

	ok, coll := database.Ingredient(ctx)
	if !ok {
		return nil, errNotConnected
	}

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(20)

	regex := bson.M{"$regex": primitive.Regex{Pattern: "\\b" + startsWith, Options: "i"}}
	startsWithBson := bson.M{"name": regex}

	ch, err := coll.Find(ctx, startsWithBson, findOptions, &dfdmodels.Ingredient{})
	if err != nil {
		return make([]*dfdmodels.Ingredient, 0), err
	}

	results := make([]*dfdmodels.Ingredient, 0)

	for i := range ch {
		results = append(results, i.(*dfdmodels.Ingredient))
	}

	return results, nil
}
