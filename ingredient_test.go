package fridgedoorapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetIngredientTree(t *testing.T) {
	ctx := context.Background()
	parentID := "5dac430246ba29343620c1df"
	parentObjectID, _ := primitive.ObjectIDFromHex(parentID)
	r, err := GetIngredientTree(ctx, parentObjectID)

	assert.Nil(t, err)
	assert.NotNil(t, r)
}
