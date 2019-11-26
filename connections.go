package fridgedoorapi

import (
	"errors"

	"github.com/digitalfridgedoor/fridgedoordatabase"
	"github.com/digitalfridgedoor/fridgedoordatabase/ingredient"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
	"github.com/digitalfridgedoor/fridgedoordatabase/user"
)

var connection fridgedoordatabase.Connection

// Connected is true if connected to mongodb
func Connected() bool {
	return connection != nil
}

// Recipe connection
func Recipe() (*recipe.Collection, error) {
	if !Connected() {
		connected := Connect()
		if !connected {
			return nil, errors.New("Cannot connect to mongodb")
		}
	}

	return recipe.New(connection), nil
}

// User connection
func User() (*user.Collection, error) {
	if !Connected() {
		connected := Connect()
		if !connected {
			return nil, errors.New("Cannot connect to mongodb")
		}
	}

	return user.New(connection), nil
}

// Ingredient connection
func Ingredient() (*ingredient.Collection, error) {
	if !Connected() {
		connected := Connect()
		if !connected {
			return nil, errors.New("Cannot connect to mongodb")
		}
	}

	return ingredient.New(connection), nil
}
