package fridgedoorapi

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoordatabase/ingredient"
)

// IngredientTreeNode represents a node in the ingredient tree
type IngredientTreeNode struct {
	ID    primitive.ObjectID    `json:"id" bson:"_id,omitempty"`
	Name  string                `json:"name"`
	Nodes []*IngredientTreeNode `json:"nodes"`
}

// SearchIngredients retrieves the ingredients matching the query
func SearchIngredients(startsWith string) ([]*ingredient.Ingredient, error) {
	i, err := Ingredient()
	if err != nil {
		return nil, err
	}

	return i.FindByName(context.Background(), startsWith)
}

// GetIngredientTree recursively gets ingredients starting with given node ID
func GetIngredientTree(ctx context.Context, parentID primitive.ObjectID) ([]*IngredientTreeNode, error) {
	i, err := Ingredient()
	if err != nil {
		return nil, err
	}

	ings := i.IngredientByParentID(ctx, parentID)
	nodes := make([]*IngredientTreeNode, 0)

	for _, ing := range ings {
		subNodes, err := GetIngredientTree(ctx, ing.ID)
		if err == nil {
			node := &IngredientTreeNode{
				ID:    ing.ID,
				Name:  ing.Name,
				Nodes: subNodes,
			}
			nodes = append(nodes, node)
		}
	}

	return nodes, nil
}
