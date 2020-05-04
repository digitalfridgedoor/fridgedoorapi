package recipeapi

import "errors"

var errAuth = errors.New("Unauthorized")
var errInvalidID = errors.New("Invalid ID")
var errNotConnected = errors.New("Not connected")
var errSubRecipe = errors.New("Invalid subrecipe reference")
var errDuplicate = errors.New("Duplicate")
