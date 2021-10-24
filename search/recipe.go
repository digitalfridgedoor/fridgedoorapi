package search

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindRecipeByName finds recipes starting with the given letter
func FindRecipeByName(ctx context.Context, startsWith string, userID primitive.ObjectID, limit int64) ([]*RecipeDescription, error) {

	ok, coll := database.Recipe(ctx)
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
		return []*RecipeDescription{}, err
	}

	results := readRecipeDescriptionFromChannel(ch, userID)
	return results, nil
}

// FindRecipeByTags finds recipes with the given tags
func FindRecipeByTags(ctx context.Context, userID primitive.ObjectID, tags []string, notTags []string, limit int64) ([]*RecipeDescription, error) {

	// https://stackoverflow.com/questions/6940503/mongodb-get-documents-by-tags

	ok, coll := database.Recipe(ctx)
	if !ok {
		return nil, errNotConnected
	}

	if limit > 20 {
		limit = 20
	}

	findOptions := options.Find()
	findOptions.SetLimit(limit)

	// { $and: [ {tags: { $all: ["tag"] } }, { tags: { $nin: ["anothertag"] } } ] }

	addedByBson := bson.M{"addedby": userID}
	andBson := []bson.M{addedByBson}

	if tags != nil && len(tags) > 0 {
		allBson := bson.M{"$all": tags}
		tagsBson := bson.M{"metadata.tags": allBson}
		andBson = append(andBson, tagsBson)
	}

	if notTags != nil && len(notTags) > 0 {
		ninBson := bson.M{"$nin": notTags}
		ninTagsBson := bson.M{"metadata.tags": ninBson}
		andBson = append(andBson, ninTagsBson)
	}

	ch, err := coll.Find(ctx, bson.M{"$and": andBson}, findOptions, &dfdmodels.Recipe{})
	if err != nil {
		return []*RecipeDescription{}, err
	}

	results := readRecipeDescriptionFromChannel(ch, userID)
	return results, nil
}

// FindPublicRecipes gets a users public recipes
func FindPublicRecipes(ctx context.Context, userID primitive.ObjectID, limit int64) ([]*RecipeDescription, error) {

	ok, coll := database.Recipe(ctx)
	if !ok {
		return nil, errNotConnected
	}

	if limit > 20 {
		limit = 20
	}

	findOptions := options.Find()
	findOptions.SetLimit(limit)

	addedByBson := bson.M{"addedby": userID}
	viewableByEveryone := bson.M{"metadata.viewableby.everyone": true}
	andBson := []bson.M{addedByBson, viewableByEveryone}

	ch, err := coll.Find(ctx, bson.M{"$and": andBson}, findOptions, &dfdmodels.Recipe{})
	if err != nil {
		return make([]*RecipeDescription, 0), err
	}

	results := readRecipeDescriptionFromChannel(ch, userID)
	return results, nil
}

func readRecipeDescriptionFromChannel(ch <-chan interface{}, userID primitive.ObjectID) []*RecipeDescription {
	results := make([]*RecipeDescription, 0)

	for i := range ch {
		r := i.(*dfdmodels.Recipe)

		if recipeapi.CanView(r, userID) {
			results = append(results, &RecipeDescription{
				ID:    r.ID,
				Name:  r.Name,
				Image: r.Metadata.Image,
			})
		}
	}

	return results
}
