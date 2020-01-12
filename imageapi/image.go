package imageapi

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Connect connects to s3
func Connect() (*s3.S3, error) {
	region := "eu-west-2"
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	return svc, nil
}

// CreatePresignedURL creates a presigned url for the requested resource
func CreatePresignedURL(svc *s3.S3, verb string, key string) (string, error) {
	req, err := createRequest(svc, "get", key)
	if err != nil {
		return "", err
	}

	return req.Presign(15 * time.Minute)
}

func createRequest(svc *s3.S3, verb string, key string) (*request.Request, error) {
	if verb == "put" {
		req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
			Bucket: aws.String("digitalfridgedoorphotos"),
			Key:    aws.String(key),
		})

		return req, nil
	} else if verb == "get" {
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String("digitalfridgedoorphotos"),
			Key:    aws.String(key),
		})

		return req, nil
	}

	return nil, errors.New("Invalid type")
}
