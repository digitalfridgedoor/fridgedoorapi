package dfdtesting

import (
	"fmt"
	"regexp"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindIngredientByNameTestPredicate can be used with SetFindFilter for searching ingredients by name
func FindIngredientByNameTestPredicate(gs *dfdmodels.Ingredient, m primitive.M) bool {
	nameval := m["name"].(bson.M)
	regexval := nameval["$regex"].(primitive.Regex)

	r := regexp.MustCompile(regexval.Pattern)

	return r.MatchString(gs.Name)
}

// SetUserViewFindByUsernamePredicate overrides logic for find users by username
func SetUserViewFindByUsernamePredicate() {
	SetUserViewFindPredicate(func(uv *dfdmodels.UserView, m primitive.M) bool {
		return m["username"] == uv.Username
	})
}

// SetRecipeFindByNameOrTagsPredicate overrides logic for find recipe
func SetRecipeFindByNameOrTagsPredicate() {
	SetRecipeFindPredicate(findRecipeByNameOrTagsTestPredicate)
}

func findRecipeByNameOrTagsTestPredicate(r *dfdmodels.Recipe, m bson.M) bool {

	contains := func(tags []string, tag string) bool {
		for _, t := range tags {
			if t == tag {
				return true
			}
		}
		return false
	}

	// first is found second is valid
	checkForAllTags := func(location bson.M) (bool, bool) {
		if t, foundTags := location["metadata.tags"]; foundTags {
			if all, ok := t.(bson.M)["$all"]; ok {
				tags := all.([]string)
				for _, t := range tags {
					if !contains(r.Metadata.Tags, t) {
						return true, false
					}
				}

				return true, true
			}
		}

		return false, false
	}

	// first is found second is valid
	checkForNotTags := func(location bson.M) (bool, bool) {
		if t, foundTags := location["metadata.tags"]; foundTags {
			if all, ok := t.(bson.M)["$nin"]; ok {
				tags := all.([]string)
				for _, t := range tags {
					if contains(r.Metadata.Tags, t) {
						return true, false
					}
				}

				return true, true
			}
		}

		return false, false
	}

	// first is found second is valid
	checkForNameRegex := func(location bson.M) (bool, bool) {
		if name, foundName := location["name"]; foundName {
		
			nameval := name.(bson.M)
			regexval := nameval["$regex"].(primitive.Regex)

			reg := regexp.MustCompile(regexval.Pattern)

			match := reg.MatchString(r.Name)
			return true, match
		}
		return false, false
	}

	// true if found and valid, false if not found or not valid
	checkForAnyKnown := func(location bson.M) (bool) {
		if found, valid := checkForNameRegex(location); found {
			return valid
		} else if found, valid := checkForAllTags(location); found {
			return valid
		} else if found, valid := checkForNotTags(location); found {
			return valid
		} 
		fmt.Println("unexpected value found")
		return false
	}

	andval := m["$and"].([]bson.M)
	addedby := (andval)[0]["addedby"].(primitive.ObjectID)

	if addedby != r.AddedBy {
		return false
	}

	if len(andval) > 1 {
		if valid := checkForAnyKnown(andval[1]); !valid {
			return false
		}
	}

	if len(andval) > 2 {
		if valid := checkForAnyKnown(andval[2]); !valid {
			return false
		}
	}

	if len(andval) > 3 {
		if valid := checkForAnyKnown(andval[3]); !valid {
			return false
		}
	}

	return true
}
