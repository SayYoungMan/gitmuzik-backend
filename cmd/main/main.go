package main

import "github.com/SayYoungMan/gitmuzik-backend/internal/receiver"

const (
	dailyPlaylistID string = "PLXzLX2ct6ysab-Gy0b1Xrm9Ka-Pg-yqmR"
)

func main() {
	receiver.ReceiveAndSavePlaylistItems(dailyPlaylistID, "test.json")

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
}
