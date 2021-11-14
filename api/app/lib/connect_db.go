package lib

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func ConnectDB() *dynamo.DB {
	sess := session.Must(session.NewSession())
	return dynamo.New(sess,
		&aws.Config{
			Region:      aws.String(os.Getenv("MY_AWS_DEFAULT_REGION")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("MY_AWS_ACCESS_KEY_ID"), os.Getenv("MY_AWS_SECRET_ACCESS_KEY"), ""),
			Endpoint:    aws.String(os.Getenv("AWS_ENDPOINT_URL")),
		})
}
