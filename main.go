package fridgedoorapi

import (
	"context"
	"fmt"

	"github.com/digitalfridgedoor/fridgedoordatabase"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Connect connects to the database using parameter from Systems Manager parameter store
func Connect() bool {
	connectionString, err := getConnectionString()

	if err != nil {
		fmt.Printf("Error getting connection string, %v.\n", err)
		return false
	}
	fmt.Printf("Got connection string: len=%v\n", len(connectionString))

	fmt.Printf("Connecting...\n")
	connection = fridgedoordatabase.Connect(context.Background(), connectionString)
	fmt.Printf("Connected.\n")

	return true
}

// Disconnect disconnects if there is a connection
func Disconnect() {
	if Connected() {
		connection.Disconnect()
		connection = nil
	}
}

func getConnectionString() (string, error) {
	region := "eu-west-2"
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return "", err
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))
	keyname := "mongodb"
	withDecryption := true

	fmt.Println("getting parameter")

	paramOutput, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           &keyname,
		WithDecryption: &withDecryption,
	})

	fmt.Println("success")

	if err != nil {
		return "", err
	}

	return *paramOutput.Parameter.Value, nil
}
