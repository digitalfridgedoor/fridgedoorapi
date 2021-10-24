package dfdtesting

import (
	"fmt"
	"regexp"

	"fridgedoorapi/dfdmodels"

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

// FindPlanByMonthAndYearTestPredicate can be used with SetFindFilter for get or create plan
func FindPlanByMonthAndYearTestPredicate(p *dfdmodels.Plan, m bson.M) bool {
	month := m["month"].(int)
	year := m["year"].(int)
	userid := m["userid"].(primitive.ObjectID)

	if p.UserID == nil {
		return false
	}

	return month == p.Month && year == p.Year && userid == *p.UserID
}

// FindPlanByMonthAndYearForGroupTestPredicate can be used with SetFindFilter for get or create plan
func FindPlanByMonthAndYearForGroupTestPredicate(p *dfdmodels.Plan, m bson.M) bool {
	month := m["month"].(int)
	year := m["year"].(int)
	planningGroupID := m["planninggroupid"].(primitive.ObjectID)

	if p.PlanningGroupID == nil {
		return false
	}

	return month == p.Month && year == p.Year && planningGroupID == *p.PlanningGroupID
}

// SetUserViewFindByUsernamePredicate overrides logic for find users by username
func SetUserViewFindByUsernamePredicate() {
	SetUserViewFindPredicate(func(uv *dfdmodels.UserView, m primitive.M) bool {
		return m["username"] == uv.Username
	})
}

// SetRecipeFindByNamePredicate overrides logic for find recipe in FindByName method
func SetRecipeFindByNamePredicate() {
	SetRecipeFindPredicate(findRecipeByNameTestPredicate)
}

// SetRecipeFindByTagsPredicate overrides logic for find recipe in FindByTags method
func SetRecipeFindByTagsPredicate() {
	SetRecipeFindPredicate(findRecipeByTagsTestPredicate)
}

func findRecipeByNameTestPredicate(gs *dfdmodels.Recipe, m primitive.M) bool {

	andval := m["$and"].([]bson.M)

	addedby := andval[1]["addedby"].(primitive.ObjectID)
	if addedby != gs.AddedBy {
		return false
	}

	nameval := andval[0]["name"].(bson.M)
	regexval := nameval["$regex"].(primitive.Regex)

	r := regexp.MustCompile(regexval.Pattern)

	return r.MatchString(gs.Name)
}

func findRecipeByTagsTestPredicate(r *dfdmodels.Recipe, m bson.M) bool {

	contains := func(tags []string, tag string) bool {
		for _, t := range tags {
			if t == tag {
				return true
			}
		}
		return false
	}

	andval := m["$and"].([]bson.M)
	addedby := (andval)[0]["addedby"].(primitive.ObjectID)

	if addedby != r.AddedBy {
		return false
	}

	if len(andval) > 1 {
		t := andval[1]["metadata.tags"].(bson.M)
		if all, ok := t["$all"]; ok {
			tags := all.([]string)
			for _, t := range tags {
				if !contains(r.Metadata.Tags, t) {
					return false
				}
			}
		} else if all, ok := t["$nin"]; ok {
			tags := all.([]string)
			for _, t := range tags {
				if contains(r.Metadata.Tags, t) {
					return false
				}
			}
		} else {
			fmt.Println("unexpected value")
			return false
		}
	}

	if len(andval) > 2 {
		t := andval[2]["metadata.tags"].(bson.M)
		if all, ok := t["$nin"]; ok {
			tags := all.([]string)
			for _, t := range tags {
				if contains(r.Metadata.Tags, t) {
					return false
				}
			}
		}
	}

	return true
}

// SetPlanningGroupFindByUser sets SetFindFilter to find planning groups for a given user
func SetPlanningGroupFindByUser() {
	SetPlanningGroupFindPredicate(findPlanningGroupForUserPredicate)
}

func findPlanningGroupForUserPredicate(gs *dfdmodels.PlanningGroup, m primitive.M) bool {
	userids := m["userids"].(bson.M)
	useridarr := userids["$in"].([]primitive.ObjectID)

	userid := useridarr[0]

	for _, id := range gs.UserIDs {
		if id == userid {
			return true
		}
	}

	return false
}

// SetClippingByNamePredicate overrides logic for find clipping in search method
func SetClippingByNamePredicate() {
	SetClippingFindPredicate(findClippingByNameTestPredicate)
}

func findClippingByNameTestPredicate(gs *dfdmodels.Clipping, m primitive.M) bool {

	andval := m["$and"].([]bson.M)

	userid := andval[1]["userid"].(primitive.ObjectID)
	if userid != gs.UserID {
		return false
	}

	nameval := andval[0]["name"].(bson.M)
	regexval := nameval["$regex"].(primitive.Regex)

	r := regexp.MustCompile(regexval.Pattern)

	return r.MatchString(gs.Name)
}
