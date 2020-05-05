package dfdtesting

import (
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoordatabase/database"

	"go.mongodb.org/mongo-driver/bson"
)

var overrides = make(map[string]*TestCollection)

// SetTestCollectionOverride sets a the database package to use a TestCollection
func SetTestCollectionOverride() {
	database.SetOverride(overrideDb)
}

// SetUserViewFindPredicate overrides the logic to get the result for Find
func SetUserViewFindPredicate(predicate func(*dfdmodels.UserView, bson.M) bool) bool {
	fn := func(value interface{}, filter bson.M) bool {
		uv := value.(*dfdmodels.UserView)
		return predicate(uv, filter)
	}

	coll := getOrAddTestCollection("recipeapi", "userviews")
	coll.findPredicate = fn
	return true
}

// SetRecipeFindPredicate overrides the logic to get the result for Find
func SetRecipeFindPredicate(predicate func(*dfdmodels.Recipe, bson.M) bool) bool {
	fn := func(value interface{}, filter bson.M) bool {
		uv := value.(*dfdmodels.Recipe)
		return predicate(uv, filter)
	}

	coll := getOrAddTestCollection("recipeapi", "recipes")
	coll.findPredicate = fn
	return true
}

// SetIngredientFindPredicate overrides the logic to get the result for Find
func SetIngredientFindPredicate(predicate func(*dfdmodels.Ingredient, bson.M) bool) bool {
	fn := func(value interface{}, filter bson.M) bool {
		uv := value.(*dfdmodels.Ingredient)
		return predicate(uv, filter)
	}

	coll := getOrAddTestCollection("recipeapi", "ingredients")
	coll.findPredicate = fn
	return true
}

// SetPlanFindPredicate overrides the logic to get the result for Find
func SetPlanFindPredicate(predicate func(*dfdmodels.Plan, bson.M) bool) bool {
	fn := func(value interface{}, filter bson.M) bool {
		uv := value.(*dfdmodels.Plan)
		return predicate(uv, filter)
	}

	coll := getOrAddTestCollection("recipeapi", "plans")
	coll.findPredicate = fn
	return true
}

func overrideDb(database string, collection string) database.ICollection {
	return getOrAddTestCollection(database, collection)
}

func getOrAddTestCollection(database string, collection string) *TestCollection {
	key := database + "_" + collection
	if val, ok := overrides[key]; ok {
		return val
	}
	overrides[key] = CreateTestCollection()
	return overrides[key]
}
