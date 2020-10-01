package clippingapi

import (
	"context"
	"errors"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Update applies the given updates to the recipeless meal
func Update(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, clippingID *primitive.ObjectID, updates map[string]string) (*dfdmodels.Clipping, error) {

	ok, coll := database.Clipping(ctx)
	if !ok {
		return nil, errNotConnected
	}

	meal, err := findClippingByID(ctx, coll, clippingID)
	if err != nil {
		return nil, err
	}

	if update, ok := updates["name"]; ok {
		meal.Name = update
	}

	err = coll.UpdateByID(ctx, clippingID, meal)

	if err != nil {
		return nil, err
	}

	return findClippingByID(ctx, coll, clippingID)
}

// AddLink adds the given link to the meal
func AddLink(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, clippingID *primitive.ObjectID, name string, url string) (*dfdmodels.Clipping, error) {

	return update(ctx, clippingID, func(clipping *dfdmodels.Clipping) error {
		hasValue := false

		for _, v := range clipping.Links {
			if v.URL == url {
				hasValue = true
			}
		}

		if !hasValue {
			clipping.Links = append(clipping.Links, dfdmodels.ClippingLink{URL: url, Name: name})
		}

		return nil
	})
}

// UpdateLink updates a link at a given index
func UpdateLink(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, clippingID *primitive.ObjectID, linkIdx int, updateProperty string, updateValue string) (*dfdmodels.Clipping, error) {

	return update(ctx, clippingID, func(clipping *dfdmodels.Clipping) error {

		if len(clipping.Links) < linkIdx {
			return errors.New("Out of range")
		}

		if updateProperty == "name" {
			clipping.Links[linkIdx].Name = updateValue
		} else if updateProperty == "url" {
			clipping.Links[linkIdx].URL = updateValue
		}

		return nil
	})
}

// SwapLinkPosition swaps the links
func SwapLinkPosition(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, clippingID *primitive.ObjectID, linkIdx1 int, linkIdx2 int) (*dfdmodels.Clipping, error) {

	return update(ctx, clippingID, func(clipping *dfdmodels.Clipping) error {

		if len(clipping.Links) < linkIdx1 || len(clipping.Links) < linkIdx2 {
			return errors.New("Out of range")
		}

		first := clipping.Links[linkIdx1]
		clipping.Links[linkIdx1] = clipping.Links[linkIdx2]
		clipping.Links[linkIdx2] = first

		return nil
	})
}

// RemoveLink removes the link at the given index
func RemoveLink(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, clippingID *primitive.ObjectID, linkIdx int) (*dfdmodels.Clipping, error) {

	return update(ctx, clippingID, func(clipping *dfdmodels.Clipping) error {

		if len(clipping.Links) < linkIdx {
			return errors.New("Out of range")
		}

		clipping.Links = append(clipping.Links[:linkIdx], clipping.Links[linkIdx+1:]...)

		return nil
	})
}

func update(ctx context.Context, id *primitive.ObjectID, update func(*dfdmodels.Clipping) error) (*dfdmodels.Clipping, error) {
	ok, coll := database.Clipping(ctx)
	if !ok {
		return nil, errNotConnected
	}

	clipping, err := findClippingByID(ctx, coll, id)
	if err != nil {
		return nil, err
	}

	if err = update(clipping); err != nil {
		return nil, err
	}

	err = coll.UpdateByID(ctx, id, clipping)

	if err != nil {
		return nil, err
	}

	return findClippingByID(ctx, coll, id)
}
