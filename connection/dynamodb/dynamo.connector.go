package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func DynamoConnection() (*dynamodb.Client, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == dynamodb.ServiceID && region == "us-west-2" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           "http://localhost:8000",
				SigningRegion: "us-west-2",
			}, nil
		}
		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		// I don't want to fallback on default resolution, i will return nil instead.
		//return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		return aws.Endpoint{}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "local", SecretAccessKey: "local", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		return nil, err
	}
	return dynamodb.NewFromConfig(cfg), nil
}
