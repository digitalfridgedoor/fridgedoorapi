package fridgedoorapi

import (
	"errors"

	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

// Recipe connection
func Recipe() (*recipe.Connection, error) {
	if !Connected() {
		connected := Connect()
		if !connected {
			return nil, errors.New("Cannot connect to mongodb")
		}
	}

	return recipe.New(connection), nil
}
