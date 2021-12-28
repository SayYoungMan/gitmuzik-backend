package connector

import (
	"context"

	"github.com/SayYoungMan/gitmuzik-backend/internal/logger"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

// Connect to S3 bucket and return the client
func ConnectToS3(ctx context.Context) *s3.S3 {
	// Initialize a session that the SDK will use to load
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create S3 client
	svc := s3.New(sess)

	logger.FromContext(ctx).Info("S3 client successfully created")

	return svc
}

func GetS3Uploader(ctx context.Context, svc *s3.S3) *s3manager.Uploader {
	// Create an uploader with the s3 Client and return
	uploader := s3manager.NewUploaderWithClient(svc)

	logger.FromContext(ctx).Info("S3 Uploader successfully created")

	return uploader
}
