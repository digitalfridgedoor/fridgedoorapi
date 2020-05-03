package fridgedoorgateway

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/aws/aws-lambda-go/events"
)

var errMissingParameter = errors.New("Parameter is missing")
var errParameterInvalid = errors.New("Parameter is an invalid")

// ReadParameterAsObjectID attempts to get an object id paramter from the path
func ReadParameterAsObjectID(request *events.APIGatewayProxyRequest, name string) (*primitive.ObjectID, error) {
	p, ok := request.PathParameters[name]
	if !ok || p == "" {
		return nil, errMissingParameter
	}

	id, err := primitive.ObjectIDFromHex(p)
	if err != nil {
		return nil, errMissingParameter
	}

	return &id, nil
}
