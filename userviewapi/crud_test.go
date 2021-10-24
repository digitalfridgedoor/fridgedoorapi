package userviewapi

import (
	"context"
	"testing"
	"time"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestUsernameCanOnlyBeUsedOnce(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()

	duration, _ := time.ParseDuration("10s")
	ctx, cancelFunc := context.WithTimeout(context.Background(), duration)
	defer cancelFunc()

	username := "userviewapi.TestUsernameCanOnlyBeUsedOnce"

	view, err := Create(ctx, username)
	assert.NotNil(t, view)
	assert.Nil(t, err)

	view, err = Create(ctx, username)
	assert.NotNil(t, err)
	assert.Equal(t, errUserExists, err)

	err = delete(ctx, username)
	assert.Nil(t, err)
}
