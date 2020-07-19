package clippingapi

import (
	"context"
	"errors"

	"github.com/digitalfridgedoor/fridgedoorapi/database"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/maisiesadler/gomongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Update applies the given updates to the recipeless meal
func Update(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, mealID *primitive.ObjectID, updates map[string]string) (*dfdmodels.Clipping, error) {

	ok, coll := database.Clipping(ctx)
	if !ok {
		return nil, errNotConnected
	}

	meal, err := findClippingByID(ctx, coll, mealID)
	if err != nil {
		return nil, err
	}

	if update, ok := updates["rename"]; ok {
		meal.Name = update
	}

	err = coll.UpdateByID(ctx, mealID, meal)

	if err != nil {
		return nil, err
	}

	return findClippingByID(ctx, coll, mealID)
}

// AddLink adds the given link to the meal
func AddLink(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, mealID *primitive.ObjectID, name string, url string) (*dfdmodels.Clipping, error) {

	return update(ctx, mealID, func(meal *dfdmodels.Clipping) error {
		hasValue := false

		for _, v := range meal.Links {
			if v.URL == url {
				hasValue = true
			}
		}

		if !hasValue {
			meal.Links = append(meal.Links, dfdmodels.Link{URL: url, Name: name})
		}

		return nil
	})
}

// RemoveLink removes the link at the given index
func RemoveLink(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, mealID *primitive.ObjectID, linkIdx int) (*dfdmodels.Clipping, error) {

	return update(ctx, mealID, func(meal *dfdmodels.Clipping) error {

		if len(meal.Links) < linkIdx {
			return errors.New("Out of range")
		}

		meal.Links = append(meal.Links[:linkIdx], meal.Links[linkIdx+1:]...)

		return nil
	})
}

func update(ctx context.Context, id *primitive.ObjectID, update func(*dfdmodels.Clipping) error) (*dfdmodels.Clipping, error) {
	ok, coll := database.Clipping(ctx)
	if !ok {
		return nil, errNotConnected
	}

	meal, err := findClippingByID(ctx, coll, id)
	if err != nil {
		return nil, err
	}

	if err = update(meal); err != nil {
		return nil, err
	}

	err = coll.UpdateByID(ctx, id, meal)

	if err != nil {
		return nil, err
	}

	return findClippingByID(ctx, coll, id)
}

func findClippingByID(ctx context.Context, coll gomongo.ICollection, id *primitive.ObjectID) (*dfdmodels.Clipping, error) {

	obj, err := coll.FindByID(ctx, id, &dfdmodels.Clipping{})
	if err != nil {
		return nil, err
	}

	meal := obj.(*dfdmodels.Clipping)

	return meal, nil
}
