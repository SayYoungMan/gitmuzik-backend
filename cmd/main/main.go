package main

import (
	"fmt"

	"github.com/SayYoungMan/gitmuzik-backend/internal/connector"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	dailyPlaylistID string = "PLXzLX2ct6ysab-Gy0b1Xrm9Ka-Pg-yqmR"
)

func main() {
	// receiver.ReceiveAndSavePlaylistItems(dailyPlaylistID, "test.json")
	client := connector.ConnectToDB()

	input := &dynamodb.ListTablesInput{}
	tables, err := client.ListTables(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	for _, n := range tables.TableNames {
		fmt.Println(*n)
	}
}
