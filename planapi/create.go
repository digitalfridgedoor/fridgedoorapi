package planapi

import (
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func create(userID primitive.ObjectID, month int, year int) (bool, *dfdmodels.Plan) {
	ok, dayLength := days(month, year)
	if !ok {
		return false, nil
	}

	days := make([]dfdmodels.Day, dayLength)
	return true, &dfdmodels.Plan{
		UserID: userID,
		Month:  month,
		Year:   year,
		Days:   days,
	}
}

func days(month int, year int) (bool, int) {
	if month > 12 || month < 1 {
		return false, 0
	}
	if year < 2000 {
		return false, 0
	}

	switch month {
	case 4:
		return true, 30
	case 6:
		return true, 30
	case 9:
		return true, 30
	case 11:
		return true, 30
	case 2:
		if year%4 == 0 {
			return true, 29
		}
		return true, 28
	default:
		return true, 31
	}
}
