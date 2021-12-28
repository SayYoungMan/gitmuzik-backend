package main

import (
	"github.com/SayYoungMan/gitmuzik-backend/internal/connector"
	"github.com/SayYoungMan/gitmuzik-backend/internal/jobticker"
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
	client := connector.ConnectToS3(ctx)
	uploader := connector.GetS3Uploader(ctx, client)
	jt := jobticker.NewJobTicker(ctx)

	for {
		<-jt.Timer.C
		receiver.ReceiveAndSavePlaylistItems(ctx, dailyPlaylistID, testFilePath)
		err := receiver.MoveFileToS3(ctx, uploader, testFilePath, testBucketName, receiver.MakeFileName(dailyPlaylistID), true)
		if err != nil {
			logger.FromContext(ctx).Fatalf("Failed to Move File to S3: %v", err)
		} else {
			logger.FromContext(ctx).Info("Successfully Moved file to S3")
		}
		jt.UpdateJobTicker(ctx)
	}
}
