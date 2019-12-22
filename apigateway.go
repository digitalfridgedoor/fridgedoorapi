package fridgedoorapi

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// ParseUsername attempts to parse the username from the Authorizer
func ParseUsername(request *events.APIGatewayProxyRequest) (string, bool) {
	if claims, ok := request.RequestContext.Authorizer["claims"]; ok {
		c := claims.(map[string]interface{})
		username, ok := c["cognito:username"]
		return username.(string), ok
	}

	return "", false
}

// ParseNickname attempts to parse the username from the Authorizer
func ParseNickname(request *events.APIGatewayProxyRequest) (string, bool) {
	if claims, ok := request.RequestContext.Authorizer["claims"]; ok {
		c := claims.(map[string]interface{})
		for k, v := range c {
			fmt.Printf("Claims has value: '%v' = '%v'.\n", k, v)
		}
		nickname, ok := c["cognito:nickname"]
		if ok {
			fmt.Printf("Got nickname: %v.\n", nickname)
			return nickname.(string), true
		}

		fmt.Printf("Could not find nickname.\n")
		return "", false
	}

	return "", false
}

// ResponseSuccessful returns a 200 response for API Gateway that allows cors
func ResponseSuccessful(body string) events.APIGatewayProxyResponse {
	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	resp.Headers["Access-Control-Allow-Headers"] = "Content-Type,Authorization,dfd-auth"
	resp.Body = body
	resp.StatusCode = 200
	return resp
}
