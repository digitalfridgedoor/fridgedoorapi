package dfdtesting

import (
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"

	"github.com/maisiesadler/gomongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var overrides = make(map[string]*TestCollection)

// SetTestCollectionOverride sets a the database package to use a TestCollection
func SetTestCollectionOverride() {
	gomongo.SetOverride(overrideDb)
	setAllIDSetters()
}

func setAllIDSetters() {
	setUserViewIDSetter()
	setRecipeIDSetter()
	setIngredientIDSetter()
}

// SetUserViewFindPredicate overrides the logic to get the result for Find
func SetUserViewFindPredicate(predicate func(*dfdmodels.UserView, bson.M) bool) bool {
	fn := func(value interface{}, filter bson.M) bool {
		uv := value.(*dfdmodels.UserView)
		return predicate(uv, filter)
	}

	coll := getOrAddTestCollection("recipeapi", "userviews")
	coll.SetFindFilter(fn)

	return true
}

func setUserViewIDSetter() {
	coll := getOrAddTestCollection("recipeapi", "userviews")
	coll.SetIDSetter(func(document interface{}, id primitive.ObjectID) {
		if u, ok := document.(*dfdmodels.UserView); ok {
			u.ID = &id
		}
	})
}

// SetRecipeFindPredicate overrides the logic to get the result for Find
func SetRecipeFindPredicate(predicate func(*dfdmodels.Recipe, bson.M) bool) bool {
	fn := func(value interface{}, filter bson.M) bool {
		uv := value.(*dfdmodels.Recipe)
		return predicate(uv, filter)
	}

	coll := getOrAddTestCollection("recipeapi", "recipes")
	coll.SetFindFilter(fn)
	return true
}

func setRecipeIDSetter() {
	coll := getOrAddTestCollection("recipeapi", "recipes")
	coll.SetIDSetter(func(document interface{}, id primitive.ObjectID) {
		if u, ok := document.(*dfdmodels.Recipe); ok {
			u.ID = &id
		}
	})
}

// SetIngredientFindPredicate overrides the logic to get the result for Find
func SetIngredientFindPredicate(predicate func(*dfdmodels.Ingredient, bson.M) bool) bool {
	fn := func(value interface{}, filter bson.M) bool {
		uv := value.(*dfdmodels.Ingredient)
		return predicate(uv, filter)
	}

	coll := getOrAddTestCollection("recipeapi", "ingredients")
	coll.SetFindFilter(fn)
	return true
}

func setIngredientIDSetter() {
	coll := getOrAddTestCollection("recipeapi", "ingredients")
	coll.SetIDSetter(func(document interface{}, id primitive.ObjectID) {
		if u, ok := document.(*dfdmodels.Ingredient); ok {
			u.ID = &id
		}
	})
}

// SetTestFindPredicate shows an example of how to override find functionality
func SetTestFindPredicate(predicate func(*dfdmodels.UserView, bson.M) bool) bool {
	fn := func(value interface{}, filter bson.M) bool {
		uv := value.(*dfdmodels.UserView)
		return predicate(uv, filter)
	}

	coll := getOrAddTestCollection("_database", "_collection")
	coll.findPredicate = fn
	return true
}

func overrideDb(database string, collection string) gomongo.ICollection {
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
