package main

import (
	"github.com/SayYoungMan/gitmuzik-backend/internal/connector"
	"github.com/SayYoungMan/gitmuzik-backend/internal/logger"
	"github.com/SayYoungMan/gitmuzik-backend/internal/receiver"
)

const (
	dailyPlaylistID string = "PLXzLX2ct6ysab-Gy0b1Xrm9Ka-Pg-yqmR"
	testFilePath    string = "test.json"
	testBucketName  string = "gitmuzik-bucket"
)

func main() {
	ctx := logger.GetNewContextWithLogger()

	// receiver.ReceiveAndSavePlaylistItems(dailyPlaylistID, "test.json")

	// Check connection to db
	// client := connector.ConnectToDB()

	// input := &dynamodb.ListTablesInput{}
	// tables, err := client.ListTables(input)
	// if err != nil {
	// 	if aerr, ok := err.(awserr.Error); ok {
	// 		switch aerr.Code() {
	// 		case dynamodb.ErrCodeInternalServerError:
	// 			fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
	// 		default:
	// 			fmt.Println(aerr.Error())
	// 		}
	// 	} else {
	// 		// Print the error, cast err to awserr.Error to get the Code and
	// 		// Message from an error.
	// 		fmt.Println(err.Error())
	// 	}
	// 	return
	// }

	// for _, n := range tables.TableNames {
	// 	fmt.Println(*n)
	// }

	// Check connection to s3
	// client := connector.ConnectToS3()
	// res, err := client.ListBuckets(nil)
	// if err != nil {
	// 	fmt.Printf("Unable to list buckets, %v", err)
	// }
	// for _, b := range res.Buckets {
	// 	fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	// }

	// Check s3 File upload
	receiver.ReceiveAndSavePlaylistItems(ctx, dailyPlaylistID, testFilePath)
	client := connector.ConnectToS3(ctx)
	uploader := connector.GetS3Uploader(ctx, client)
	err := receiver.MoveFileToS3(ctx, uploader, testFilePath, testBucketName, "test-key", true)
	if err != nil {
		logger.FromContext(ctx).Fatalf("Failed to Move File to S3: %v", err)
	} else {
		logger.FromContext(ctx).Info("Successfully Moved file to S3")
	}
}
