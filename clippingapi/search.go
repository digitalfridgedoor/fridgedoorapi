package clippingapi

import (
	"context"
	"digitalfridgedoor/fridgedoorapi/database"
	"digitalfridgedoor/fridgedoorapi/dfdmodels"

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
	addedByBson := bson.M{"addedby": userID}
	andBson := bson.M{"$and": []bson.M{startsWithBson, addedByBson}}

	ch, err := coll.Find(ctx, andBson, findOptions, &dfdmodels.Recipe{})
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
			ID:   r.ID,
			Name: r.Name,
		})
	}

	return results
}
