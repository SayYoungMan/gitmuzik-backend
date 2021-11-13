package connector

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Connect to AWS DynamoDB and return the client
func ConnectToDB() *dynamodb.DynamoDB {
	// Initialize a session that the SDK will use to load
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create Dynamo DB client
	svc := dynamodb.New(sess)

	return svc
}
