package userviewapi

import "errors"

var errUserExists = errors.New("User exists")
var errNotConnected = errors.New("Not connected")
var errNotUpdated = errors.New("Not updated")
