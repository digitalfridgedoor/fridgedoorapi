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

	return FindRecipe(ctx, userID, startsWith, []string{}, []string{}, limit)
}

// FindRecipeByTags finds recipes with the given tags
func FindRecipeByTags(ctx context.Context, userID primitive.ObjectID, tags []string, notTags []string, limit int64) ([]*RecipeDescription, error) {

	return FindRecipe(ctx, userID, "", tags, notTags, limit)
}

// FindRecipe finds recipes by name or tags
func FindRecipe(ctx context.Context, userID primitive.ObjectID, startsWith string, tags []string, notTags []string, limit int64) ([]*RecipeDescription, error) {
  // { $and: [ {"name": {$regex: "\\bF"}}, {"metadata.tags": { $all: ["weeknight"] }} ] }
  // { $and: [ {"metadata.tags": { $all: ["tag"] } }, { "metadata.tags": { $nin: ["anothertag"] } } ] }

  andBson := []bson.M{}

  andBson = appendAddedByBson(andBson, userID)
  andBson = appendStartsWithBson(andBson, startsWith)
  andBson = appendTags(andBson, tags)
  andBson = appendNotTags(andBson, notTags)

  return findAndBson(ctx, andBson, userID, limit)	
}

// FindPublicRecipes gets a users public recipes
func FindPublicRecipes(ctx context.Context, userID primitive.ObjectID, limit int64) ([]*RecipeDescription, error) {

	addedByBson := bson.M{"addedby": userID}
	viewableByEveryone := bson.M{"metadata.viewableby.everyone": true}
	andBson := []bson.M{addedByBson, viewableByEveryone}

	return findAndBson(ctx, andBson, userID, limit)	
}

func appendAddedByBson(andBson []primitive.M, userID primitive.ObjectID) ([]primitive.M) {
	addedByBson := bson.M{"addedby": userID}
	return append(andBson, addedByBson)
}

func appendStartsWithBson(andBson []primitive.M, startsWith string) ([]primitive.M) {
	if len(startsWith) > 0 {
		regex := bson.M{"$regex": primitive.Regex{Pattern: "\\b" + startsWith, Options: "i"}}
		startsWithBson := bson.M{"name": regex}
		return append(andBson, startsWithBson)
	}
	return andBson
}

func appendTags(andBson []primitive.M, tags []string) ([]primitive.M) {
	// https://stackoverflow.com/questions/6940503/mongodb-get-documents-by-tags

	if tags != nil && len(tags) > 0 {
		allBson := bson.M{"$all": tags}
		tagsBson := bson.M{"metadata.tags": allBson}
		return append(andBson, tagsBson)
	}

	return andBson
}

func appendNotTags(andBson []primitive.M, notTags []string) ([]primitive.M) {
	if notTags != nil && len(notTags) > 0 {
		ninBson := bson.M{"$nin": notTags}
		ninTagsBson := bson.M{"metadata.tags": ninBson}
		return append(andBson, ninTagsBson)
	}

	return andBson
}

func findAndBson(ctx context.Context, andBson []bson.M, userID primitive.ObjectID, limit int64) ([]*RecipeDescription, error) {

	ok, coll := database.Recipe(ctx)
	if !ok {
		return nil, errNotConnected
	}

	if limit > 20 {
		limit = 20
	}

	findOptions := options.Find()
	findOptions.SetLimit(limit)

	ch, err := coll.Find(ctx, bson.M{"$and": andBson}, findOptions, &dfdmodels.Recipe{})
	if err != nil {
		return []*RecipeDescription{}, err
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
