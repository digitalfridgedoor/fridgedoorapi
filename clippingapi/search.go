package clippingapi

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SearchByName finds clippings starting with the given letter
func SearchByName(ctx context.Context, startsWith string, userID primitive.ObjectID, limit int64) ([]*ClippingDescription, error) {

	ok, coll := database.Clipping(ctx)
	if !ok {
		return nil, errNotConnected
	}

	if limit > 20 {
		limit = 20
	}

	findOptions := options.Find()
	findOptions.SetLimit(limit)

	regex := bson.M{"$regex": primitive.Regex{Pattern: "\\b" + startsWith, Options: "i"}}
	startsWithBson := bson.M{"name": regex}
	useridBson := bson.M{"userid": userID}
	andBson := bson.M{"$and": []bson.M{startsWithBson, useridBson}}

	ch, err := coll.Find(ctx, andBson, findOptions, &dfdmodels.Clipping{})
	if err != nil {
		return []*ClippingDescription{}, err
	}

	results := readClippingViewFromChannel(ch, userID)
	return results, nil
}

func readClippingViewFromChannel(ch <-chan interface{}, userID primitive.ObjectID) []*ClippingDescription {
	results := make([]*ClippingDescription, 0)

	for i := range ch {
		r := i.(*dfdmodels.Clipping)

		results = append(results, &ClippingDescription{
			ID:       r.ID,
			Name:     r.Name,
			RecipeID: r.RecipeID,
		})
	}

	return results
}
